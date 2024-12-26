package internal 

import (
	"testing"

	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestBTPostOrder(t *testing.T) {
	tests := []struct {
		name     string
		tree     *in.BinaryNode[int]
		expected []int
	}{
		{
			name: "Complete binary tree",
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
			expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "Single node",
			tree: &in.BinaryNode[int]{
				Value: 42,
			},
			expected: []int{42},
		},
		{
			name:     "Empty tree",
			tree:     nil,
			expected: []int{},
		},
		{
			name: "Left-skewed tree",
			tree: &in.BinaryNode[int]{
				Value: 1,
				Left: &in.BinaryNode[int]{
					Value: 2,
					Left: &in.BinaryNode[int]{
						Value: 3,
					},
				},
			},
			expected: []int{3, 2, 1},
		},
		{
			name: "Right-skewed tree",
			tree: &in.BinaryNode[int]{
				Value: 1,
				Right: &in.BinaryNode[int]{
					Value: 2,
					Right: &in.BinaryNode[int]{
						Value: 3,
					},
				},
			},
			expected: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.BTPostOrder(tt.tree)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Expected %v, but got %v", tt.expected, result)
					break
				}
			}
		})
	}
}
