package bst

// Node represents a node of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// New returns a new node with key `key` and no children
func New(key int) *Node {
	return &Node{Key: key}
}

// Map walks through tree nodes in order, applying function `f` to every node
func (n *Node) Map(f func(*Node)) {
	if n.Left != nil {
		n.Left.Map(f)
	}
	f(n)
	if n.Right != nil {
		n.Right.Map(f)
	}
}

// Equals compares two trees by value (structural equality)
func Equals(n1 *Node, n2 *Node) bool {
	eq := true
	if n1.Left != nil && n2.Left != nil {
		eq = eq && Equals(n1.Left, n2.Left)
	} else if (n1.Left == nil && n2.Left != nil) || (n1.Left != nil && n2.Left == nil) {
		return false
	}
	if n1.Key != n2.Key {
		return false
	}
	if n1.Right != nil && n2.Right != nil {
		eq = eq && Equals(n1.Right, n2.Right)
	} else if (n1.Right == nil && n2.Right != nil) || (n1.Right != nil && n2.Right == nil) {
		return false
	}
	return eq
}

// Keys returns a slice of all tree keys in order
func (n *Node) Keys() []int {
	keys := []int{}
	n.Map(func(node *Node) {
		keys = append(keys, node.Key)
	})
	return keys
}

// Insert inserts key `key` to subtree, preserving order, duplicates are dropped silently
func (n *Node) Insert(key int) {
	for {
		if key < n.Key {
			if n.Left != nil {
				n = n.Left
			} else {
				n.Left = &Node{key, nil, nil}
				break
			}
		} else if key > n.Key {
			if n.Right != nil {
				n = n.Right
			} else {
				n.Right = &Node{key, nil, nil}
				break
			}
		} else {
			break
		}
	}
}

// Find finds node by key
func (n *Node) Find(key int) *Node {
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
func (n *Node) Min() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// Max finds node with the maximum key (the rightmost one in BST)
func (n *Node) Max() *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

// Remove removes node with key `key` from tree (does nothing if not found)
func (n *Node) Remove(key int) {
	for n != nil && n.Key != key {
		if key < n.Key {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	if n == nil {
		return
	}

}
