package main

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid [][]int
		res  int
	}{
		{[][]int{}, 0},
		{[][]int{{1}}, 1},
		{[][]int{{0, 1}}, 1},
		{[][]int{{0}, {1}}, 1},
		{[][]int{{0, 1}, {1, 0}}, 1},

		{[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
	}
	for _, test := range tests {
		maxArea := maxAreaOfIsland(test.grid)
		if maxArea != test.res {
			t.Errorf("failed: got %d,want %d\n", maxArea, test.res)
			break
		}
	}
}
