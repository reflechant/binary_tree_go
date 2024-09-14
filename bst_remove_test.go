package bst

import (
	"slices"
	"testing"
)

func TestRemoveNonExistent(t *testing.T) {
	n1 := &Node[int, string]{1, "",
		&Node[int, string]{-2, "", nil, nil},
		&Node[int, string]{2, "", nil,
			&Node[int, string]{3, "", nil, nil}}}
	n2 := &Node[int, string]{1, "",
		&Node[int, string]{-2, "", nil, nil}, &Node[int, string]{2, "", nil,
			&Node[int, string]{3, "", nil, nil}}}

	n1 = n1.Remove(7)

	if !n1.Left.EqualKeys(n2) {
		t.FailNow()
	}
}

func TestRemoveRoot(t *testing.T) {
	n1 := &Node[int, string]{2, "",
		&Node[int, string]{1, "", nil, nil},
		&Node[int, string]{3, "", nil, nil}}
	n2 := &Node[int, string]{3, "",
		&Node[int, string]{1, "", nil, nil}, nil}

	n1 = n1.Remove(2)

	if !n1.EqualKeys(n2) {
		t.FailNow()
	}
}

func TestRemoveGeneral(t *testing.T) {
	n1 := new(Node[int, string])
	n1.Insert(5, "")
	n1.Insert(3, "")
	n1.Insert(7, "")
	n1.Insert(2, "")
	n1.Insert(8, "")
	n1.Insert(1, "")

	n1 = n1.Remove(5)

	if !slices.Equal(
		slices.Collect(n1.Keys()),
		[]int{1, 2, 3, 7, 8},
	) {
		t.FailNow()
	}
}
