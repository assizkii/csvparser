package entities

import (
	"sync"
)

type Place struct {
	Id        int
	Name      string
	Condition string
	State     string
	Price     string
	PriceInt  int
}

type BTree struct {
	Root      *node
	height    int
	MaxHeight int
	mx        sync.Mutex
}

type node struct {
	left  *node
	right *node
	data  Place
}

func (t *BTree) Insert(data Place) {
	t.mx.Lock()
	defer t.mx.Unlock()

	if t.Root == nil {
		t.Root = &node{data: data}
	} else {
		if t.height >= t.MaxHeight {
			if data.PriceInt >= t.Root.data.PriceInt {
				return
			}
		}
		t.Root.insert(data)
		t.height++
	}
}
//insert node to btree
func (n *node) insert(data Place) {
	if data.PriceInt <= n.data.PriceInt {
		if n.left == nil {
			n.left = &node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}
// btree to array
func (n *node) ToArray(result []Place) []Place {
	if n == nil {
		return nil
	}
	if n.left != nil {
		result = n.left.ToArray(result)
	}
	result = append(result, n.data)
	if n.right != nil {
		result = n.right.ToArray(result)
	}
	return result
}
