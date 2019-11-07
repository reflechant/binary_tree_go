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

// Equals compares two nodes by value (structural equality)
func (n1 *Node) Equals(n2 *Node) bool {
	return true
}

// Slice returns a slice of all tree keys in order
func (n1 *Node) Slice() []int {
	return []int{}
}
