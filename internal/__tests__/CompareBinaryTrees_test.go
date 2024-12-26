package internal

import (
	"testing"

	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestCompareBinaryTrees(t *testing.T) {
	tests := []struct {
		name     string
		treeA    *in.BinaryNode[int]
		treeB    *in.BinaryNode[int]
		expected bool
	}{
		{
			name: "Identical trees",
			treeA: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 5,
				},
				Right: &in.BinaryNode[int]{
					Value: 15,
				},
			},
			treeB: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 5,
				},
				Right: &in.BinaryNode[int]{
					Value: 15,
				},
			},
			expected: true,
		},
		{
			name: "Different tree structure",
			treeA: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 5,
				},
			},
			treeB: &in.BinaryNode[int]{
				Value: 10,
				Right: &in.BinaryNode[int]{
					Value: 15,
				},
			},
			expected: false,
		},
		{
			name: "Different values in nodes",
			treeA: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 5,
				},
				Right: &in.BinaryNode[int]{
					Value: 15,
				},
			},
			treeB: &in.BinaryNode[int]{
				Value: 10,
				Left: &in.BinaryNode[int]{
					Value: 6,
				},
				Right: &in.BinaryNode[int]{
					Value: 15,
				},
			},
			expected: false,
		},
		{
			name:     "Empty trees",
			treeA:    nil,
			treeB:    nil,
			expected: true,
		},
		{
			name: "One tree is empty",
			treeA: &in.BinaryNode[int]{
				Value: 10,
			},
			treeB: nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.CompareBinaryTrees(tt.treeA, tt.treeB)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
