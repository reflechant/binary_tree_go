package bst

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestOneNodeMinMax(t *testing.T) {
	n1 := Node[int, string]{1, "", nil, nil}
	min := n1.Min()
	assert.Equal(t, 1, min.Key)
	max := n1.Max()
	assert.Equal(t, 1, max.Key)
}

func TestTwoLeavesMinMax(t *testing.T) {
	n1 := Node[int, string]{2, "", &Node[int, string]{1, "", nil, nil}, &Node[int, string]{3, "", nil, nil}}
	min := n1.Min()
	assert.Equal(t, 1, min.Key)
	max := n1.Max()
	assert.Equal(t, 3, max.Key)
}
