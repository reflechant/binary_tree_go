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
	keys1, keys2 := n1.Slice(), n2.Slice()
	if len(keys1) != len(keys2) {
		return false
	}
	for i := 0; i < len(keys1); i++ {
		if keys1[i] != keys2[i] {
			return false
		}
	}
	return true
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
