package internal

import (
	"testing"

	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestBTBFS(t *testing.T) {
	tests := []struct {
		name     string
		tree     *in.BinaryNode[int]
		needle   int
		expected bool
	}{
		{
			name: "Value exists in the tree",
			tree: &in.BinaryNode[int]{
				Value: 1,
				Left: &in.BinaryNode[int]{
					Value: 2,
					Left: &in.BinaryNode[int]{
						Value: 4,
					},
					Right: &in.BinaryNode[int]{
						Value: 5,
					},
				},
				Right: &in.BinaryNode[int]{
					Value: 3,
					Left: &in.BinaryNode[int]{
						Value: 6,
					},
					Right: &in.BinaryNode[int]{
						Value: 7,
					},
				},
			},
			needle:   5,
			expected: true,
		},
		{
			name: "Value does not exist in the tree",
			tree: &in.BinaryNode[int]{
				Value: 1,
				Left: &in.BinaryNode[int]{
					Value: 2,
				},
				Right: &in.BinaryNode[int]{
					Value: 3,
				},
			},
			needle:   10,
			expected: false,
		},
		{
			name: "Empty tree",
			tree: nil,
			needle:   5,
			expected: false,
		},
		{
			name: "Single node matches",
			tree: &in.BinaryNode[int]{
				Value: 42,
			},
			needle:   42,
			expected: true,
		},
		{
			name: "Single node does not match",
			tree: &in.BinaryNode[int]{
				Value: 42,
			},
			needle:   7,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.BTBFS(tt.tree, tt.needle)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
