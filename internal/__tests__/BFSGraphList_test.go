package internal

import (
	"testing"

	in "github.com/vinewz/alGOs/internal"
	alg "github.com/vinewz/alGOs/src/_today"
)

func TestBFSGraphList(t *testing.T) {
	tests := []struct {
		name     string
		graph    in.WeightedAdjacencyList
		source   int
		needle   int
		expected []int
	}{
		{
			name: "Path exists",
			graph: in.WeightedAdjacencyList{
				0: {{To: 1, Weight: 1}, {To: 2, Weight: 1}},
				1: {{To: 3, Weight: 1}},
				2: {{To: 3, Weight: 1}},
				3: {},
			},
			source:   0,
			needle:   3,
			expected: []int{0, 1, 3},
		},
		{
			name: "No path exists",
			graph: in.WeightedAdjacencyList{
				0: {{To: 1, Weight: 1}},
				1: {{To: 2, Weight: 1}},
				2: {},
				3: {},
			},
			source:   0,
			needle:   3,
			expected: nil,
		},
		{
			name: "Source is the needle",
			graph: in.WeightedAdjacencyList{
				0: {{To: 1, Weight: 1}},
				1: {{To: 2, Weight: 1}},
				2: {},
			},
			source:   0,
			needle:   0,
			expected: []int{0},
		},
		{
			name:     "Empty graph",
			graph:    in.WeightedAdjacencyList{},
			source:   0,
			needle:   0,
			expected: nil,
		},
		{
			name: "Disconnected graph",
			graph: in.WeightedAdjacencyList{
				0: {},
				1: {},
				2: {},
			},
			source:   0,
			needle:   2,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.BFSGraphList(tt.graph, tt.source, tt.needle)
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
