package bst

import "testing"

func TestAppendRightToNode(t *testing.T) {
	n1 := Node{1, nil, nil}
	n1.Append(2)
	n2 := Node{1, nil, &Node{2, nil, nil}}
	if !Equals(&n1, &n2) {
		t.FailNow()
	}
}

func TestAppendLeftToNode(t *testing.T) {
	n1 := Node{1, nil, nil}
	n1.Append(0)
	n2 := Node{1, &Node{0, nil, nil}, nil}
	if !Equals(&n1, &n2) {
		t.FailNow()
	}
}

func TestAppendDuplicateCausesNoMutation(t *testing.T) {
	n1 := Node{1, nil, nil}
	n1.Append(1)
	n2 := Node{1, nil, nil}
	if !Equals(&n1, &n2) {
		t.FailNow()
	}
}

func TestAppendLeftToTree(t *testing.T) {
	n1 := Node{1, &Node{0, nil, nil}, nil}
	n1.Append(-1)
	n2 := Node{1, &Node{0, &Node{-1, nil, nil}, nil}, nil}
	if !Equals(&n1, &n2) {
		t.FailNow()
	}
}
func TestAppendRightToTree(t *testing.T) {
	n1 := Node{1, nil, &Node{2, nil, nil}}
	n1.Append(3)
	n2 := Node{1, nil, &Node{2, nil, &Node{3, nil, nil}}}
	if !Equals(&n1, &n2) {
		t.FailNow()
	}
}
