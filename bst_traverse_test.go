package bst

import "testing"

func TestSingleNodeTraversal(t *testing.T) {
	n1 := Node{1, nil, nil}
	keys := make([]int, 0)
	n1.Traverse(
		func(n *Node) {
			keys = append(keys, n.Key)
		})
	if !(len(keys) == 1 && keys[0] == 1) {
		t.FailNow()
	}
}

func TestTraverseNodeWithLeft(t *testing.T) {
	n1 := Node{1, &Node{0, nil, nil}, nil}
	keys := make([]int, 0)
	n1.Traverse(
		func(n *Node) {
			keys = append(keys, n.Key)
		})
	if !(len(keys) == 2 && keys[0] == 0 && keys[1] == 1) {
		t.FailNow()
	}
}

func TestTraverseNodeWithRight(t *testing.T) {
	n1 := Node{1, nil, &Node{2, nil, nil}}
	keys := make([]int, 0)
	n1.Traverse(
		func(n *Node) {
			keys = append(keys, n.Key)
		})
	if !(len(keys) == 2 && keys[0] == 1 && keys[1] == 2) {
		t.FailNow()
	}
}