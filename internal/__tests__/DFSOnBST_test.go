package internal

import (
	"testing"

	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestDFSOnBST(t *testing.T) {
	tests := []struct {
		name     string
		tree     *in.BinaryNode[int]
		needle   int
		expected bool
	}{
		{
			name: "Element exists in BST",
			tree: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 5,
					Left: &in.BinaryNode[int]{
						Value: 3,
					},
					Right: &in.BinaryNode[int]{
						Value: 7,
					},
				},
				Right: &in.BinaryNode[int]{
					Value: 15,
					Left: &in.BinaryNode[int]{
						Value: 12,
					},
					Right: &in.BinaryNode[int]{
						Value: 18,
					},
				},
			},
			needle:   7,
			expected: true,
		},
		{
			name: "Element does not exist in BST",
			tree: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 5,
					Left: &in.BinaryNode[int]{
						Value: 3,
					},
					Right: &in.BinaryNode[int]{
						Value: 7,
					},
				},
				Right: &in.BinaryNode[int]{
					Value: 15,
					Left: &in.BinaryNode[int]{
						Value: 12,
					},
					Right: &in.BinaryNode[int]{
						Value: 18,
					},
				},
			},
			needle:   6,
			expected: false,
		},
		{
			name: "Element is root of BST",
			tree: &in.BinaryNode[int]{
				Value: 10,
			},
			needle:   10,
			expected: true,
		},
		{
			name: "Empty tree",
			tree: nil,
			needle:   5,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.DFSOnBST(tt.tree, tt.needle)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
