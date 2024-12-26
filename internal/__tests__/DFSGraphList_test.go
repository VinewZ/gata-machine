package internal

import (
	"testing"

	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestDFSGraphList(t *testing.T) {
	tests := []struct {
		name     string
		graph    in.WeightedAdjacencyList
		source   int
		needle   int
		expected []int
	}{
		{
			name: "Path exists in the graph",
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}, in.GraphEdge{To: 2, Weight: 2}},
				{in.GraphEdge{To: 3, Weight: 1}},
				{in.GraphEdge{To: 3, Weight: 1}},
				{},
			},
			source:   0,
			needle:   3,
			expected: []int{0, 1, 3},
		},
		{
			name: "No path exists",
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}},
				{in.GraphEdge{To: 2, Weight: 1}},
				{},
			},
			source:   0,
			needle:   3,
			expected: nil,
		},
		{
			name: "Source is the needle",
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}},
				{in.GraphEdge{To: 2, Weight: 1}},
				{},
			},
			source:   0,
			needle:   0,
			expected: []int{0},
		},
		{
			name: "Empty graph",
			graph: in.WeightedAdjacencyList{
				{},
			},
			source:   0,
			needle:   1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.DFSGraphList(tt.graph, tt.source, tt.needle)
			if !equalDFS(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func equalDFS(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
