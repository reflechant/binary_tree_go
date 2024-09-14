package bst

import "testing"

func TestFindRoot(t *testing.T) {
	n1 := Node[int,string]{1, "", nil, nil}
	n2 := n1.Find(1)
	if &n1 != n2 {
		t.FailNow()
	}
}

func TestFindLeftChild(t *testing.T) {
	n1 := Node[int,string]{1, "", &Node[int,string]{-1, "", nil, nil}, nil}
	n2 := n1.Find(-1)
	if n2 != n1.Left {
		t.FailNow()
	}
}

func TestFindRightChild(t *testing.T) {
	n1 := Node[int,string]{1, "", nil, &Node[int,string]{2, "", nil, nil}}
	n2 := n1.Find(2)
	if n2 != n1.Right {
		t.FailNow()
	}
}

func TestFindNonExistent(t *testing.T) {
	n1 := Node[int,string]{1, "", nil, &Node[int,string]{2, "", nil, nil}}
	n2 := n1.Find(5)
	if n2 != nil {
		t.FailNow()
	}
}
