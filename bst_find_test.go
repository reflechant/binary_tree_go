package bst

import "testing"

func TestFindRoot(t *testing.T) {
	n1 := Node{1, nil, nil}
	n2 := n1.Find(1)
	if &n1 != n2 {
		t.FailNow()
	}
}

func TestFindLeftChild(t *testing.T) {
	n1 := Node{1, &Node{-1, nil, nil}, nil}
	n2 := n1.Find(-2)
	if n2 != n1.Left {
		t.FailNow()
	}
}

func TestFindRightChild(t *testing.T) {
	n1 := Node{1, nil, &Node{2, nil, nil}}
	n2 := n1.Find(2)
	if n2 != n1.Left {
		t.FailNow()
	}
}
