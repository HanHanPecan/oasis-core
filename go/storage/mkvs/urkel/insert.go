package urkel

import (
	"fmt"

	"github.com/oasislabs/ekiden/go/common/crypto/hash"
	"github.com/oasislabs/ekiden/go/storage/mkvs/urkel/internal"
)

func (t *Tree) doInsert(ptr *internal.Pointer, depth uint8, key hash.Hash, val []byte) (*internal.Pointer, bool, error) {
	node, err := t.cache.derefNodePtr(internal.NodeID{Path: key, Depth: depth}, ptr, nil)
	if err != nil {
		return nil, false, err
	}

	switch n := node.(type) {
	case nil:
		// Insert into nil node, create a new leaf node.
		return t.cache.newLeafNode(key, val), false, nil
	case *internal.InternalNode:
		var existed bool
		if getKeyBit(key, depth) {
			n.Right, existed, err = t.doInsert(n.Right, depth+1, key, val)
		} else {
			n.Left, existed, err = t.doInsert(n.Left, depth+1, key, val)
		}
		if err != nil {
			return nil, false, err
		}

		if !n.Left.IsClean() || !n.Right.IsClean() {
			n.Clean = false
			ptr.Clean = false
		}

		return ptr, existed, nil
	case *internal.LeafNode:
		// If the key matches, we can just update the value.
		if n.Key.Equal(&key) {
			if n.Value.Equal(val) {
				return ptr, true, nil
			}

			t.cache.removeValue(n.Value)
			n.Value = t.cache.newValue(val)
			n.Clean = false
			ptr.Clean = false
			return ptr, true, nil
		}

		existingBit := getKeyBit(n.Key, depth)
		newBit := getKeyBit(key, depth)

		var left, right *internal.Pointer
		if existingBit != newBit {
			// No bit collision at this depth, create an internal node with
			// two leaves.
			if existingBit {
				left = t.cache.newLeafNode(key, val)
				right = ptr
			} else {
				left = ptr
				right = t.cache.newLeafNode(key, val)
			}
		} else {
			// Bit collision at this depth.
			if existingBit {
				left = nil
				right, _, err = t.doInsert(ptr, depth+1, key, val)
			} else {
				left, _, err = t.doInsert(ptr, depth+1, key, val)
				right = nil
			}
			if err != nil {
				return nil, false, err
			}
		}

		return t.cache.newInternalNode(left, right), false, nil
	default:
		panic(fmt.Sprintf("urkel: unknown node type: %+v", n))
	}
}