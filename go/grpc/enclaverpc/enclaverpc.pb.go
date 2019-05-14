// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enclaverpc/enclaverpc.proto

package enclaverpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CallEnclaveRequest struct {
	// Raw request payload that will be passed to the enclave.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// Endpoint identifier for cases where a single node supports multiple endpoints.
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallEnclaveRequest) Reset()         { *m = CallEnclaveRequest{} }
func (m *CallEnclaveRequest) String() string { return proto.CompactTextString(m) }
func (*CallEnclaveRequest) ProtoMessage()    {}
func (*CallEnclaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aad1ebbbcd47deea, []int{0}
}

func (m *CallEnclaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallEnclaveRequest.Unmarshal(m, b)
}
func (m *CallEnclaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallEnclaveRequest.Marshal(b, m, deterministic)
}
func (m *CallEnclaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallEnclaveRequest.Merge(m, src)
}
func (m *CallEnclaveRequest) XXX_Size() int {
	return xxx_messageInfo_CallEnclaveRequest.Size(m)
}
func (m *CallEnclaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallEnclaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallEnclaveRequest proto.InternalMessageInfo

func (m *CallEnclaveRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CallEnclaveRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type CallEnclaveResponse struct {
	// Raw response payload from enclave.
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallEnclaveResponse) Reset()         { *m = CallEnclaveResponse{} }
func (m *CallEnclaveResponse) String() string { return proto.CompactTextString(m) }
func (*CallEnclaveResponse) ProtoMessage()    {}
func (*CallEnclaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aad1ebbbcd47deea, []int{1}
}

func (m *CallEnclaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallEnclaveResponse.Unmarshal(m, b)
}
func (m *CallEnclaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallEnclaveResponse.Marshal(b, m, deterministic)
}
func (m *CallEnclaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallEnclaveResponse.Merge(m, src)
}
func (m *CallEnclaveResponse) XXX_Size() int {
	return xxx_messageInfo_CallEnclaveResponse.Size(m)
}
func (m *CallEnclaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallEnclaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallEnclaveResponse proto.InternalMessageInfo

func (m *CallEnclaveResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*CallEnclaveRequest)(nil), "enclaverpc.CallEnclaveRequest")
	proto.RegisterType((*CallEnclaveResponse)(nil), "enclaverpc.CallEnclaveResponse")
}

func init() { proto.RegisterFile("enclaverpc/enclaverpc.proto", fileDescriptor_aad1ebbbcd47deea) }

var fileDescriptor_aad1ebbbcd47deea = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcd, 0x4b, 0xce,
	0x49, 0x2c, 0x4b, 0x2d, 0x2a, 0x48, 0xd6, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x4a, 0x5e, 0x5c, 0x42, 0xce, 0x89, 0x39, 0x39, 0xae, 0x10, 0x91, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0x48, 0x8a, 0x8b, 0x23, 0x35, 0x2f, 0xa5,
	0x20, 0x3f, 0x33, 0xaf, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x57, 0xd2, 0xe7,
	0x12, 0x46, 0x31, 0xab, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xb7, 0x61, 0x46, 0x71, 0x5c, 0x5c,
	0x30, 0xc5, 0x05, 0xc9, 0x42, 0x01, 0x5c, 0xdc, 0x48, 0xda, 0x85, 0xe4, 0xf4, 0x90, 0x1c, 0x8e,
	0xe9, 0x46, 0x29, 0x79, 0x9c, 0xf2, 0x10, 0x7b, 0x95, 0x18, 0x9c, 0x0c, 0xa2, 0xf4, 0xd2, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xf3, 0x13, 0x8b, 0x33, 0x8b, 0x73, 0x12,
	0x93, 0x8a, 0xf5, 0x53, 0xb3, 0x33, 0x53, 0x52, 0xf3, 0xf4, 0xd3, 0xf3, 0xf5, 0xd3, 0x51, 0x03,
	0x28, 0x89, 0x0d, 0x1c, 0x42, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x6e, 0x9c, 0xbe,
	0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EnclaveRpcClient is the client API for EnclaveRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnclaveRpcClient interface {
	// RPC interface.
	CallEnclave(ctx context.Context, in *CallEnclaveRequest, opts ...grpc.CallOption) (*CallEnclaveResponse, error)
}

type enclaveRpcClient struct {
	cc *grpc.ClientConn
}

func NewEnclaveRpcClient(cc *grpc.ClientConn) EnclaveRpcClient {
	return &enclaveRpcClient{cc}
}

func (c *enclaveRpcClient) CallEnclave(ctx context.Context, in *CallEnclaveRequest, opts ...grpc.CallOption) (*CallEnclaveResponse, error) {
	out := new(CallEnclaveResponse)
	err := c.cc.Invoke(ctx, "/enclaverpc.EnclaveRpc/CallEnclave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnclaveRpcServer is the server API for EnclaveRpc service.
type EnclaveRpcServer interface {
	// RPC interface.
	CallEnclave(context.Context, *CallEnclaveRequest) (*CallEnclaveResponse, error)
}

func RegisterEnclaveRpcServer(s *grpc.Server, srv EnclaveRpcServer) {
	s.RegisterService(&_EnclaveRpc_serviceDesc, srv)
}

func _EnclaveRpc_CallEnclave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallEnclaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnclaveRpcServer).CallEnclave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enclaverpc.EnclaveRpc/CallEnclave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnclaveRpcServer).CallEnclave(ctx, req.(*CallEnclaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnclaveRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "enclaverpc.EnclaveRpc",
	HandlerType: (*EnclaveRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallEnclave",
			Handler:    _EnclaveRpc_CallEnclave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "enclaverpc/enclaverpc.proto",
}