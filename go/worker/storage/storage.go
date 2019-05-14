package storage

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/oasislabs/ekiden/go/common"
	"github.com/oasislabs/ekiden/go/common/crypto/hash"
	"github.com/oasislabs/ekiden/go/common/grpc"
	"github.com/oasislabs/ekiden/go/common/logging"
	"github.com/oasislabs/ekiden/go/common/node"
	epochtime "github.com/oasislabs/ekiden/go/epochtime/api"
	"github.com/oasislabs/ekiden/go/genesis"
	"github.com/oasislabs/ekiden/go/storage"
	storageApi "github.com/oasislabs/ekiden/go/storage/api"
	"github.com/oasislabs/ekiden/go/worker/registration"
)

const (
	cfgWorkerStorageEnabled = "worker.storage.enabled"
)

// Storage is a worker handling storage operations.
type Storage struct {
	enabled      bool
	storage      storageApi.Backend
	grpc         *grpc.Server
	initCh       chan struct{}
	quitCh       chan struct{}
	logger       *logging.Logger
	registration *registration.Registration
}

// New constructs a new storage worker.
func New(
	epochtime epochtime.Backend,
	sb storageApi.Backend,
	g *grpc.Server,
	r *registration.Registration,
	consensus common.ConsensusBackend,
	genesis genesis.Provider,
) (*Storage, error) {

	s := &Storage{
		enabled:      viper.GetBool(cfgWorkerStorageEnabled),
		storage:      sb,
		grpc:         g,
		initCh:       make(chan struct{}),
		quitCh:       make(chan struct{}),
		logger:       logging.GetLogger("worker/storage"),
		registration: r,
	}

	if s.enabled {
		// Populate storage from genesis.
		consensus.RegisterGenesisHook(func() {
			doc, err := genesis.GetGenesisDocument()
			if err != nil {
				s.logger.Error("failed to get genesis document",
					"err", err,
				)
				panic("failed to get genesis document")
			}

			if err = s.initGenesis(&doc.Storage); err != nil {
				s.logger.Error("failed to initialize storage from genesis",
					"err", err,
				)
				panic("storage: failed to initialize storage from genesis")
			}
		})

		// Attach storage worker to gRPC server.
		storage.NewGRPCServer(s.grpc.Server(), s.storage)

		// Register storage worker role.
		s.registration.RegisterRole(func(n *node.Node) error {
			n.AddRoles(node.RoleStorageWorker)

			return nil
		})
	}

	return s, nil
}

// Name returns the service name.
func (s *Storage) Name() string {
	return "storage worker"
}

// Enabled returns if worker is enabled.
func (s *Storage) Enabled() bool {
	return s.enabled
}

// Initialized returns a channel that will be closed when the storage worker
// is initialized and ready to service requests.
func (s *Storage) Initialized() <-chan struct{} {
	return s.initCh
}

// Start starts the storage service.
func (s *Storage) Start() error {
	if !s.enabled {
		s.logger.Info("not starting storage worker as it is disabled")

		// In case the worker is not enabled, close the init channel immediately.
		close(s.initCh)

		return nil
	}

	// Wait for the node to be registered for the current epoch.
	go func() {
		s.logger.Info("starting storage worker, waiting for registration")
		<-s.registration.InitialRegistrationCh()

		s.logger.Info("storage worker started")

		close(s.initCh)
	}()

	return nil
}

// Stop halts the service.
func (s *Storage) Stop() {
	close(s.quitCh)
}

// Quit returns a channel that will be closed when the service terminates.
func (s *Storage) Quit() <-chan struct{} {
	return s.quitCh
}

// Cleanup performs the service specific post-termination cleanup.
func (s *Storage) Cleanup() {
}

func (s *Storage) initGenesis(gen *storageApi.Genesis) error {
	ctx := context.Background()

	s.logger.Info("initializing storage from genesis")

	var emptyRoot hash.Hash
	emptyRoot.Empty()
	if _, err := s.storage.Apply(ctx, emptyRoot, emptyRoot, gen.State); err != nil {
		return err
	}

	return nil
}

// RegisterFlags registers the configuration flags with the provided
// command.
func RegisterFlags(cmd *cobra.Command) {
	if !cmd.Flags().Parsed() {
		cmd.Flags().Bool(cfgWorkerStorageEnabled, false, "Enable storage worker")
	}
	for _, v := range []string{
		cfgWorkerStorageEnabled,
	} {
		viper.BindPFlag(v, cmd.Flags().Lookup(v)) // nolint: errcheck
	}
}