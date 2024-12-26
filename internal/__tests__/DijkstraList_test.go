package internal

import (
	"testing"

	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestDijkstraList(t *testing.T) {
	tests := []struct {
		name     string
		source   int
		sink     int
		graph    in.WeightedAdjacencyList
		expected []int
	}{
		{
			name:   "Path exists with shortest path",
			source: 0,
			sink:   3,
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}, in.GraphEdge{To: 2, Weight: 4}},
				{in.GraphEdge{To: 2, Weight: 2}, in.GraphEdge{To: 3, Weight: 5}},
				{in.GraphEdge{To: 3, Weight: 1}},
				{},
			},
			expected: []int{0, 1, 2, 3},
		},
		{
			name:   "No path exists",
			source: 0,
			sink:   3,
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}},
				{in.GraphEdge{To: 2, Weight: 2}},
				{},
			},
			expected: nil,
		},
		{
			name:   "Source is the sink",
			source: 2,
			sink:   2,
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}},
				{in.GraphEdge{To: 2, Weight: 2}},
				{},
			},
			expected: []int{2},
		},
		{
			name:   "Graph with multiple paths",
			source: 0,
			sink:   3,
			graph: in.WeightedAdjacencyList{
				{in.GraphEdge{To: 1, Weight: 1}, in.GraphEdge{To: 2, Weight: 5}},
				{in.GraphEdge{To: 3, Weight: 1}},
				{in.GraphEdge{To: 3, Weight: 2}},
				{},
			},
			expected: []int{0, 1, 3},
		},
		{
			name:   "Empty graph",
			source: 0,
			sink:   0,
			graph:  in.WeightedAdjacencyList{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.DijkstraList(tt.source, tt.sink, tt.graph)
			if !equalDIJ(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func equalDIJ(a, b []int) bool {
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
