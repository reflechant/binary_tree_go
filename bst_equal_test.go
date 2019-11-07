package bst

import "testing"

func TestSingleNodesAreEqual(t *testing.T) {
	n1 := Node{1, nil, nil}
	n2 := Node{1, nil, nil}
	if !n1.Equals(&n2) || !n2.Equals(&n1) {
		t.FailNow()
	}
}

func TestNodesWithChildrenAreEqual(t *testing.T) {
	c1 := Node{1, nil, nil}
	c2 := Node{3, nil, nil}
	n1 := Node{2, &c1, &c2}
	n2 := Node{2, &c1, &c2}
	if !n1.Equals(&n2) || !n2.Equals(&n1) {
		t.FailNow()
	}
}

func TestDifferentNodesAreNotEqual(t *testing.T) {
	n1 := Node{1, nil, nil}
	n2 := Node{2, nil, nil}
	if n1.Equals(&n2) || n2.Equals(&n1) {
		t.FailNow()
	}
	n1 = Node{1, &Node{0, nil, nil}, nil}
	n2 = Node{1, nil, nil}
	if n1.Equals(&n2) || n2.Equals(&n1) {
		t.FailNow()
	}
}
