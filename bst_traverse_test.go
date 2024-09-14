package bst

import (
	"slices"
	"testing"
)

func TestSingleNodeTraversal(t *testing.T) {
	n1 := Node[int, string]{1, "", nil, nil}
	keys := make([]int, 0)
	n1.Map(
		func(n *Node[int, string]) {
			keys = append(keys, n.Key)
		})
		
	if !slices.Equal(keys, []int{1}) {
		t.FailNow()
	}
}

func TestTraverseNodeWithLeft(t *testing.T) {
	n1 := Node[int, string]{1, "",
		&Node[int, string]{0, "", nil, nil}, nil}
	keys := make([]int, 0)
	n1.Map(
		func(n *Node[int, string]) {
			keys = append(keys, n.Key)
		})
	if !slices.Equal(keys, []int{0, 1}) {
		t.FailNow()
	}
}

func TestTraverseNodeWithRight(t *testing.T) {
	n1 := Node[int, string]{1, "", nil,
		&Node[int, string]{2, "", nil, nil}}
	keys := make([]int, 0)
	n1.Map(
		func(n *Node[int, string]) {
			keys = append(keys, n.Key)
		})
	if !slices.Equal(keys, []int{1, 2}) {
		t.FailNow()
	}
}
