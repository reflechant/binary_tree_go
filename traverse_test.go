package bst

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSingleNodeTraversal(t *testing.T) {
	n1 := Node[int, string]{1, "", nil, nil}
	keys := make([]int, 0)
	n1.Map(
		func(n *Node[int, string]) {
			keys = append(keys, n.Key)
		})

	assert.Equal(t, []int{1}, keys)
}

func TestTraverseNodeWithLeft(t *testing.T) {
	n1 := Node[int, string]{1, "",
		&Node[int, string]{0, "", nil, nil}, nil}
	keys := make([]int, 0)
	n1.Map(
		func(n *Node[int, string]) {
			keys = append(keys, n.Key)
		})
	assert.Equal(t, []int{0, 1}, keys)
}

func TestTraverseNodeWithRight(t *testing.T) {
	n1 := Node[int, string]{1, "", nil,
		&Node[int, string]{2, "", nil, nil}}
	keys := make([]int, 0)
	n1.Map(
		func(n *Node[int, string]) {
			keys = append(keys, n.Key)
		})
	assert.Equal(t, []int{1, 2}, keys)
}
