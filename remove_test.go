package bst

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNonExistent(t *testing.T) {
	n1 := &Node[int, string]{1, "",
		&Node[int, string]{-2, "", nil, nil},
		&Node[int, string]{2, "", nil,
			&Node[int, string]{3, "", nil, nil}}}

	n2 := &Node[int, string]{1, "",
		&Node[int, string]{-2, "", nil, nil},
		&Node[int, string]{2, "", nil,
			&Node[int, string]{3, "", nil, nil}}}

	n1 = n1.Remove(7)

	assert.True(t, n1.EqualKeys(n2))
}

func TestRemoveRoot(t *testing.T) {
	n1 := &Node[int, string]{2, "",
		&Node[int, string]{1, "", nil, nil},
		&Node[int, string]{3, "", nil, nil}}
	n2 := &Node[int, string]{3, "",
		&Node[int, string]{1, "", nil, nil}, nil}

	n1 = n1.Remove(2)

	assert.True(t, n1.EqualKeys(n2))
}

func TestRemoveGeneral(t *testing.T) {
	n1 := &Node[int, string]{5, "", nil, nil}
	n1.Insert(3, "")
	n1.Insert(7, "")
	n1.Insert(2, "")
	n1.Insert(8, "")
	n1.Insert(1, "")

	n1 = n1.Remove(5)

	assert.Equal(t, []int{1, 2, 3, 7, 8}, slices.Collect(n1.Keys()))
}
