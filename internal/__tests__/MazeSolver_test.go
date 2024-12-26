package internal

import (
	"testing"
	alg "github.com/vinewz/alGOs/src/_today"
	in "github.com/vinewz/alGOs/internal"
)

func TestMazeSolver(t *testing.T) {
	tests := []struct {
		name     string
		maze     []string
		wall     string
		start    in.Point
		end      in.Point
		expected []in.Point
	}{
		{
			name: "Simple maze, path exists",
			maze: []string{
				"xxxxxxxxxx x",
				"x        x x",
				"x        x x",
				"x xxxxxxxx x",
				"x          x",
				"x xxxxxxxxxx",
			},
			wall:  "x",
			start: in.Point{X: 10, Y: 0},
			end:   in.Point{X: 1, Y: 5},
			expected: []in.Point{
				{X: 10, Y: 0},
				{X: 10, Y: 1},
				{X: 10, Y: 2},
				{X: 10, Y: 3},
				{X: 10, Y: 4},
				{X: 9, Y: 4},
				{X: 8, Y: 4},
				{X: 7, Y: 4},
				{X: 6, Y: 4},
				{X: 5, Y: 4},
				{X: 4, Y: 4},
				{X: 3, Y: 4},
				{X: 2, Y: 4},
				{X: 1, Y: 4},
				{X: 1, Y: 5},
			},
		},
		{
			name: "No path available",
			maze: []string{
				"xxxxxxxxxx x",
				"x        x x",
				"x        x x",
				"x xxxxxxxx x",
				"x          x",
				"x xxxxxxxxxx",
			},
			wall:  "x",
			start: in.Point{X: 0, Y: 0},
			end:   in.Point{X: 5, Y: 5},
			expected: nil,
		},
		{
			name: "Start is the same as the end",
			maze: []string{
				"xxxxxxxxxx x",
				"x        x x",
				"x        x x",
				"x xxxxxxxx x",
				"x          x",
				"x xxxxxxxxxx",
			},
			wall:  "x",
			start: in.Point{X: 0, Y: 0},
			end:   in.Point{X: 0, Y: 0},
			expected: []in.Point{
				{X: 0, Y: 0},
			},
		},
		{
			name: "Empty maze",
			maze: []string{},
			wall:  "x",
			start: in.Point{X: 0, Y: 0},
			end:   in.Point{X: 0, Y: 0},
			expected: nil,
		},
		{
			name: "Maze with blocked paths",
			maze: []string{
				"xxxxxxxxxx x",
				"x        x x",
				"x xxxxxx x x",
				"x          x",
				"x xxxxxxxxxx",
			},
			wall:  "x",
			start: in.Point{X: 0, Y: 0},
			end:   in.Point{X: 10, Y: 5},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := alg.MazeSolver(tt.maze, tt.wall, tt.start, tt.end)
			if !equalPoints(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func equalPoints(a, b []in.Point) bool {
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
