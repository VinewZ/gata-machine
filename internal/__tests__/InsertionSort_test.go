package internal

import (
	"testing"
	alg "github.com/vinewz/alGOs/src/_today"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sort an unsorted list",
			input:    []int{5, 2, 9, 1, 5, 6},
			expected: []int{1, 2, 5, 5, 6, 9},
		},
		{
			name:     "Already sorted list",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "List with one element",
			input:    []int{3},
			expected: []int{3},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "List with negative numbers",
			input:    []int{-5, -2, -9, 1, 5, 0},
			expected: []int{-9, -5, -2, 0, 1, 5},
		},
		{
			name:     "List with duplicates",
			input:    []int{5, 3, 5, 3, 4},
			expected: []int{3, 3, 4, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alg.InsertionSort(tt.input)
			for i, v := range tt.input {
				if v != tt.expected[i] {
					t.Errorf("At index %d, expected %d, but got %d", i, tt.expected[i], v)
				}
			}
		})
	}
}
