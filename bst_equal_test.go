package bst

import "testing"

func TestEquals(t *testing.T) {
	cases := []struct {
		name     string
		a        Node[int, string]
		b        Node[int, string]
		expected bool
	}{
		{"single nodes equality test",
			Node[int, string]{1, "a", nil, nil},
			Node[int, string]{1, "b", nil, nil},
			true},
		{"tree equality test",
			Node[int, string]{2, "",
				&Node[int, string]{1, "", nil, nil},
				&Node[int, string]{3, "", nil, nil}},
			Node[int, string]{2, "",
				&Node[int, string]{1, "", nil, nil},
				&Node[int, string]{3, "", nil, nil}},
			true},
		{"nodes with different keys are not equal",
			Node[int, string]{1, "", nil, nil},
			Node[int, string]{2, "", nil, nil},
			false,
		},
		{"single node not equals node with left child",
			Node[int, string]{1, "",
				&Node[int, string]{0, "", nil, nil}, nil},
			Node[int, string]{1, "", nil, nil},
			false,
		},
		{"single node not equals node with right child",
			Node[int, string]{1, "", nil,
				&Node[int, string]{2, "", nil, nil}},
			Node[int, string]{1, "", nil, nil},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.a.EqualKeys(&tc.b) != tc.expected ||
				tc.b.EqualKeys(&tc.a) != tc.expected {
				t.Fail()
			}
		})
	}
}
