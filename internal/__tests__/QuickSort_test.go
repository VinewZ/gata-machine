package internal

import (
	alg "github.com/vinewz/alGOs/src/_today"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element list",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Already sorted list",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted list",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "List with duplicate elements",
			input:    []int{5, 2, 8, 2, 5, 1},
			expected: []int{1, 2, 2, 5, 5, 8},
		},
		{
			name:     "List with negative numbers",
			input:    []int{-5, 2, -8, 0, 5, -1},
			expected: []int{-8, -5, -1, 0, 2, 5},
		},
		{
			name:     "List with all same elements",
			input:    []int{3, 3, 3, 3, 3},
			expected: []int{3, 3, 3, 3, 3},
		},
		{
			name: "Mix of positive, negative and zero",
			input: []int{0, -1, 5, -4, 2, 0, 7},
			expected: []int{-4, -1, 0, 0, 2, 5, 7},
		},
		{
			name: "Larger unsorted list",
			input: []int{12, 3, 1, 45, 7, 8, 2, 9, 10, 5, 100},
			expected: []int{1, 2, 3, 5, 7, 8, 9, 10, 12, 45, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)
			alg.QuickSort(inputCopy)
			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, inputCopy)
			}
		})
	}
}
