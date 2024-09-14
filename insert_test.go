package bst

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		name     string
		initial  Node[int, string]
		key      int
		value    string
		expected Node[int, string]
	}{
		{"insert right to single node",
			Node[int, string]{1, "", nil, nil}, 2, "v",
			Node[int, string]{1, "", nil, &Node[int, string]{2, "v", nil, nil}}},
		{"insert left to single node",
			Node[int, string]{1, "", nil, nil}, 0, "v",
			Node[int, string]{1, "", &Node[int, string]{0, "v", nil, nil}, nil}},
		{"insert duplicate",
			Node[int, string]{1, "", nil, nil}, 1, "v",
			Node[int, string]{1, "v", nil, nil}},
		{"insert left to tree",
			Node[int, string]{1, "", &Node[int, string]{0, "", nil, nil}, nil}, -1, "v",
			Node[int, string]{1, "", &Node[int, string]{0, "", &Node[int, string]{-1, "v", nil, nil}, nil}, nil}},
		{"insert right to tree",
			Node[int, string]{1, "", nil, &Node[int, string]{2, "", nil, nil}}, 3, "v",
			Node[int, string]{1, "", nil, &Node[int, string]{2, "", nil, &Node[int, string]{3, "v", nil, nil}}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.initial.Insert(tc.key, tc.value)
			assert.True(t, tc.initial.EqualKeys(&tc.expected))
		})
	}
}
