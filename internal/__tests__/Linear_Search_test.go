package internal

import (
	alg "github.com/vinewz/alGOs/src/_today"
	"testing"
)

func TestLinear_Search(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected bool
	}{
		{
			name:     "Found element in the list",
			input:    []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			name:     "Element not found in the list",
			input:    []int{1, 2, 3, 4, 5},
			target:   6,
			expected: false,
		},
		{
			name:     "Empty list",
			input:    []int{},
			target:   3,
			expected: false,
		},
		{
			name:     "Search for the first element",
			input:    []int{1, 2, 3, 4, 5},
			target:   1,
			expected: true,
		},
		{
			name:     "Search for the last element",
			input:    []int{1, 2, 3, 4, 5},
			target:   5,
			expected: true,
		},
		{
			name:     "Search in a list with duplicates",
			input:    []int{1, 2, 3, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			name:     "Multiple occurrences of the target",
			input:    []int{1, 1, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "Search for a negative number in a list",
			input:    []int{-5, -4, -3, -2, -1},
			target:   -3,
			expected: true,
		},
		{
			name:     "Search for a positive number in a list of negative numbers",
			input:    []int{-5, -4, -3, -2, -1},
			target:   3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.LinearSearchList(tt.input, tt.target)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
