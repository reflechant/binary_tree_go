package bst

// Node represents a node of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// NewNode returns a new node with key `key` and no children
func NewNode(key int) (n Node) {
	n.Key = key
	return
}

// Traverse walks through tree nodes in order, applying function `f` to every node
func (n *Node) Traverse(f func(*Node)) {
	if n.Left != nil {
		n.Left.Traverse(f)
	}
	f(n)
	if n.Right != nil {
		n.Right.Traverse(f)
	}
}

// Equals compares two nodes by value (structural equality)
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

// Slice returns a slice of all tree keys in order
func (n *Node) Slice() []int {
	keys := []int{}
	n.Traverse(func(node *Node) {
		keys = append(keys, node.Key)
	})
	return keys
}

// Append adds key `key` to subtree, preserving order
func (n *Node) Append(key int) {
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
