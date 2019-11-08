package bst

import "testing"

func TestSingleNodesAreEqual(t *testing.T) {
	n1 := Node{1, nil, nil}
	n2 := Node{1, nil, nil}
	if !Equals(&n1, &n2) || !Equals(&n2, &n1) {
		t.FailNow()
	}
}

func TestNodesWithChildrenAreEqual(t *testing.T) {
	c1 := Node{1, nil, nil}
	c2 := Node{3, nil, nil}
	n1 := Node{2, &c1, &c2}
	n2 := Node{2, &c1, &c2}
	if !Equals(&n1, &n2) || !Equals(&n2, &n1) {
		t.FailNow()
	}
}

func TestNodesWithDifferentKeysAreNotEqual(t *testing.T) {
	n1 := Node{1, nil, nil}
	n2 := Node{2, nil, nil}
	if !Equals(&n1, &n2) || !Equals(&n2, &n1) {
		t.FailNow()
	}
}

func TestNodeWithChildrenNotEqualsNodeWithoutChildren(t *testing.T) {
	n1 := Node{1, &Node{0, nil, nil}, nil}
	n2 := Node{1, nil, nil}
	if !Equals(&n1, &n2) || !Equals(&n2, &n1) {
		t.FailNow()
	}
}
