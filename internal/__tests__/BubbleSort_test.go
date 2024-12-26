package internal

import (
	"testing"
  alg "github.com/vinewz/alGOs/src/_today"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{3, 1, 4, 5, 2},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Array with duplicates",
			input:    []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alg.BubbleSort(tt.input)
			for i := range tt.input {
				if tt.input[i] != tt.expected[i] {
					t.Errorf("Expected %v, but got %v", tt.expected, tt.input)
					break
				}
			}
		})
	}
}
