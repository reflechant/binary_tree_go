package bst

// Node represents a node of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// NewNode returns a new node with key `key` and no children
func NewNode(key int) (n *Node) {
	n.Key = key
	return
}

// Traverse walks through tree nodes in order, applying function `f` to every node
func (n *Node) Traverse(f func(node *Node)) {
	if n.Left != nil {
		n.Left.Traverse(f)
	}
	f(n)
	if n.Right != nil {
		n.Right.Traverse(f)
	}
}

// Equals compares two nodes by value (structural equality)
func Equals(n1, n2 *Node) bool {
	return true
}

// Slice returns a slice of all tree keys in order
func (n *Node) Slice() []int {
	return []int{}
}
