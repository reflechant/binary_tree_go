package bst

import (
	"cmp"
	"iter"
)

// Node represents a generic binary search tree
type Node[K cmp.Ordered, V any] struct {
	Key   K
	Value V
	Left  *Node[K, V]
	Right *Node[K, V]
}

// All returns an iterator of all tree (key,value) pairs in order
func (n *Node[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if n.Left != nil {
			n.Left.All()
		}
		if !yield(n.Key, n.Value) {
			return
		}
		if n.Right != nil {
			n.Right.All()
		}
	}
}

// Keys returns an iterator of all tree keys in order
func (n *Node[K, V]) Keys2() iter.Seq[K] {
	return func(yield func(K) bool) {
		if n.Left != nil {
			for k := range n.Left.Keys() {
				yield(k)
			}
		}
		if !yield(n.Key) {
			return
		}
		if n.Right != nil {
			for k := range n.Right.Keys() {
				yield(k)
			}
		}
	}
}

// Keys returns an iterator of all tree keys in order
func (n *Node[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		if n.Left != nil {
			n.Left.keys(yield)
		}
		if !yield(n.Key) {
			return
		}
		if n.Right != nil {
			n.Right.keys(yield)
		}
	}
}

func (n *Node[K, V]) keys(yield func(K) bool) {
	if n.Left != nil {
		n.Left.keys(yield)
	}
	if !yield(n.Key) {
		return
	}
	if n.Right != nil {
		n.Right.keys(yield)
	}
}

// Values returns an iterator of all tree values in order
func (n *Node[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		if n.Left != nil {
			n.Left.Values()
		}
		if !yield(n.Value) {
			return
		}
		if n.Right != nil {
			n.Right.Values()
		}
	}
}

// Map walks through tree nodes in order, applying function `f` to every node and returning an iterator over the resulting tree
func (n *Node[K, V]) Map(f func(*Node[K, V])) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if n.Left != nil {
			n.Left.Map(f)
		}
		f(n)
		if !yield(n.Key, n.Value) {
			return
		}
		if n.Right != nil {
			n.Right.Map(f)
		}
	}
}

// EqualKeys compares trees by key
func (n *Node[K, V]) EqualKeys(n2 *Node[K, V]) bool {
	next1, stop1 := iter.Pull(n.Keys())
	defer stop1()
	next2, stop2 := iter.Pull(n2.Keys())
	defer stop2()

	k1, ok1 := next1()
	k2, ok2 := next2()
	for ok1 || ok2 {
		// one sequence is shorter -> unequal
		if ok1 != ok2 {
			return false
		}

		if k1 != k2 {
			return false
		}

		k1, ok1 = next1()
		k2, ok2 = next2()
	}

	return true
}

// Insert inserts a (key,value) pair to the tree, preserving order, duplicates are dropped silently
func (n *Node[K, V]) Insert(key K, value V) {
	for {
		if key < n.Key {
			if n.Left != nil {
				n = n.Left
			} else {
				n.Left = &Node[K, V]{key, value, nil, nil}
				break
			}
		} else if key > n.Key {
			if n.Right != nil {
				n = n.Right
			} else {
				n.Right = &Node[K, V]{key, value, nil, nil}
				break
			}
		} else {
			break
		}
	}
}

// Find finds node by key
func (n *Node[K, V]) Find(key K) *Node[K, V] {
	for n != nil {
		if key == n.Key {
			return n
		} else if key < n.Key {
			n = n.Left
		} else if key > n.Key {
			n = n.Right
		}
	}
	return nil
}

// Min finds node with the minimum key (the leftmost one in BST)
func (n *Node[K, V]) Min() *Node[K, V] {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// Max finds node with the maximum key (the rightmost one in BST)
func (n *Node[K, V]) Max() *Node[K, V] {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

// Remove removes node with key `key` from tree (does nothing if not found)
func (n *Node[K, V]) Remove(key K) *Node[K, V] {
	if n == nil {
		return nil
	}
	if key < n.Key {
		n.Left = n.Left.Remove(key)
	} else if key > n.Key {
		n.Right = n.Right.Remove(key)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}
		minRight := n.Right.Min()
		n.Key = minRight.Key
		n.Right = n.Right.Remove(minRight.Key)
	}
	return n
}
