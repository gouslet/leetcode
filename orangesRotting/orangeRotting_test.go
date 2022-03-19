/*
 * File: \orangesRotting\orangeRotting_test.go                                 *
 * Project: leetcode                                                           *
 * Created At: Wednesday, 2022/03/16 , 11:42:17                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 17:25:56                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "testing"

func TestOrangeRotting(t *testing.T) {
	tests := []struct {
		grid    [][]int
		maxTime int
	}{
		{[][]int{{}}, -1},
		{[][]int{{1}}, -1},
		{[][]int{{2}}, 0},
		{[][]int{{0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{1, 2}}, 1},
		{[][]int{{0}, {1}}, -1},
		{[][]int{{0}, {2}}, 0},
		{[][]int{{1}, {2}}, 1},
		{[][]int{{0, 0, 0}}, 0},
		{[][]int{{0, 0, 1}}, -1},
		{[][]int{{0, 0, 2}}, 0},
		{[][]int{{0, 1, 0}}, -1},
		{[][]int{{0, 1, 1}}, -1},
		{[][]int{{0, 1, 2}}, 1},
		{[][]int{{0, 2, 0}}, 0},
		{[][]int{{0, 2, 1}}, 1},
		{[][]int{{1, 0, 0}}, -1},
		{[][]int{{1, 0, 1}}, -1},
		{[][]int{{1, 0, 2}}, -1},
		{[][]int{{1, 1, 0}}, -1},
		{[][]int{{1, 1, 1}}, -1},
		{[][]int{{1, 1, 2}}, 2},
		{[][]int{{1, 2, 0}}, 1},
		{[][]int{{1, 2, 1}}, 1},
		{[][]int{{1, 2, 2}}, 1},
		{[][]int{{2, 0, 0}}, 0},
		{[][]int{{2, 0, 1}}, -1},
		{[][]int{{2, 0, 2}}, 0},
		{[][]int{{2, 1, 0}}, 1},
		{[][]int{{2, 1, 1}}, 2},
		{[][]int{{2, 1, 2}}, 1},
		{[][]int{{2, 2, 0}}, 0},
		{[][]int{{2, 2, 1}}, 1},
		{[][]int{{2, 2, 2}}, 0},

		{[][]int{{0}, {0}, {0}}, 0},
		{[][]int{{0}, {0}, {1}}, -1},
		{[][]int{{0}, {0}, {2}}, 0},
		{[][]int{{0}, {1}, {0}}, -1},
		{[][]int{{0}, {1}, {1}}, -1},
		{[][]int{{0}, {1}, {2}}, 1},
		{[][]int{{0}, {2}, {0}}, 0},
		{[][]int{{0}, {2}, {1}}, 1},
		{[][]int{{1}, {0}, {0}}, -1},
		{[][]int{{1}, {0}, {1}}, -1},
		{[][]int{{1}, {0}, {2}}, -1},
		{[][]int{{1}, {1}, {0}}, -1},
		{[][]int{{1}, {1}, {1}}, -1},
		{[][]int{{1}, {1}, {2}}, 2},
		{[][]int{{1}, {2}, {0}}, 1},
		{[][]int{{1}, {2}, {1}}, 1},
		{[][]int{{1}, {2}, {2}}, 1},
		{[][]int{{2}, {0}, {0}}, 0},
		{[][]int{{2}, {0}, {1}}, -1},
		{[][]int{{2}, {0}, {2}}, 0},
		{[][]int{{2}, {1}, {0}}, 1},
		{[][]int{{2}, {1}, {1}}, 2},
		{[][]int{{2}, {1}, {2}}, 1},
		{[][]int{{2}, {2}, {0}}, 0},
		{[][]int{{2}, {2}, {1}}, 1},
		{[][]int{{2}, {2}, {2}}, 0},

		{[][]int{{0, 0}, {0, 0}}, 0},
		{[][]int{{0, 0}, {0, 1}}, -1},
		{[][]int{{0, 0}, {0, 2}}, 0},
		{[][]int{{0, 0}, {1, 0}}, -1},
		{[][]int{{0, 0}, {1, 1}}, -1},
		{[][]int{{0, 0}, {1, 2}}, 1},
		{[][]int{{0, 0}, {2, 0}}, 0},
		{[][]int{{0, 0}, {2, 1}}, 1},
		{[][]int{{0, 0}, {2, 2}}, 0},
		{[][]int{{0, 1}, {0, 0}}, -1},
		{[][]int{{0, 1}, {0, 1}}, -1},
		{[][]int{{0, 1}, {1, 0}}, -1},
		{[][]int{{0, 1}, {1, 1}}, -1},
		{[][]int{{0, 1}, {1, 2}}, 1},
		{[][]int{{0, 1}, {2, 0}}, -1},
		{[][]int{{0, 1}, {2, 1}}, 2},
		{[][]int{{0, 1}, {2, 2}}, 1},
		{[][]int{{0, 2}, {0, 0}}, 0},
		{[][]int{{0, 2}, {0, 1}}, 1},
		{[][]int{{0, 2}, {0, 2}}, 0},
		{[][]int{{0, 2}, {1, 0}}, -1},
		{[][]int{{0, 2}, {1, 1}}, 2},
		{[][]int{{0, 2}, {1, 2}}, 1},
		{[][]int{{0, 2}, {2, 0}}, 0},
		{[][]int{{0, 2}, {2, 1}}, 1},
		{[][]int{{0, 2}, {2, 2}}, 0},
		{[][]int{{1, 0}, {0, 0}}, -1},
		{[][]int{{1, 0}, {0, 1}}, -1},
		{[][]int{{1, 0}, {0, 2}}, -1},
		{[][]int{{1, 0}, {1, 0}}, -1},
		{[][]int{{1, 0}, {1, 1}}, -1},
		{[][]int{{1, 0}, {1, 2}}, 2},
		{[][]int{{1, 0}, {2, 0}}, 1},
		{[][]int{{1, 0}, {2, 1}}, 1},
		{[][]int{{1, 0}, {2, 2}}, 1},
		{[][]int{{1, 1}, {0, 0}}, -1},
		{[][]int{{1, 1}, {0, 1}}, -1},
		{[][]int{{1, 1}, {0, 2}}, 2},
		{[][]int{{1, 1}, {1, 0}}, -1},
		{[][]int{{1, 1}, {1, 1}}, -1},
		{[][]int{{1, 1}, {1, 2}}, 2},
		{[][]int{{1, 1}, {2, 0}}, 2},
		{[][]int{{1, 1}, {2, 1}}, 2},
		{[][]int{{1, 1}, {2, 2}}, 1},
		{[][]int{{1, 2}, {0, 0}}, 1},
		{[][]int{{1, 2}, {0, 1}}, 1},
		{[][]int{{1, 2}, {0, 2}}, 1},
		{[][]int{{1, 2}, {1, 0}}, 2},
		{[][]int{{1, 2}, {1, 1}}, 2},
		{[][]int{{1, 2}, {1, 2}}, 1},
		{[][]int{{1, 2}, {2, 0}}, 1},
		{[][]int{{1, 2}, {2, 1}}, 1},
		{[][]int{{1, 2}, {2, 2}}, 1},
		{[][]int{{2, 0}, {0, 0}}, 0},
		{[][]int{{2, 0}, {0, 1}}, -1},
		{[][]int{{2, 0}, {0, 2}}, 0},
		{[][]int{{2, 0}, {1, 0}}, 1},
		{[][]int{{2, 0}, {1, 1}}, 2},
		{[][]int{{2, 0}, {1, 2}}, 1},
		{[][]int{{2, 0}, {2, 0}}, 0},
		{[][]int{{2, 0}, {2, 1}}, 1},
		{[][]int{{2, 0}, {2, 2}}, 0},
		{[][]int{{2, 1}, {0, 0}}, 1},
		{[][]int{{2, 1}, {0, 1}}, 2},
		{[][]int{{2, 1}, {0, 2}}, 1},
		{[][]int{{2, 1}, {1, 0}}, 1},
		{[][]int{{2, 1}, {1, 1}}, 2},
		{[][]int{{2, 1}, {1, 2}}, 1},
		{[][]int{{2, 1}, {2, 0}}, 1},
		{[][]int{{2, 1}, {2, 1}}, 1},
		{[][]int{{2, 1}, {2, 2}}, 1},
		{[][]int{{2, 2}, {0, 0}}, 0},
		{[][]int{{2, 2}, {0, 1}}, 1},
		{[][]int{{2, 2}, {0, 2}}, 0},
		{[][]int{{2, 2}, {1, 0}}, 1},
		{[][]int{{2, 2}, {1, 1}}, 1},
		{[][]int{{2, 2}, {1, 2}}, 1},
		{[][]int{{2, 2}, {2, 0}}, 0},
		{[][]int{{2, 2}, {2, 1}}, 1},
		{[][]int{{2, 2}, {2, 2}}, 0},

		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, -1},
		{[][]int{{2, 0, 0}, {0, 1, 1}, {0, 0, 0}, {0, 1, 1}}, -1},
		{[][]int{{2, 0, 1}, {0, 1, 1}, {0, 2, 0}, {0, 1, 1}}, 3},
	}
	for _, test := range tests {
		// printMatrix := func(matrix [][]int) {
		// 	for _, line := range matrix {
		// 		t.Logf("%v\n", line)
		// 	}
		// }
		var ori_grid [][]int = make([][]int, len(test.grid))
		for i := range ori_grid {
			ori_grid[i] = make([]int, len(test.grid[0]))
			copy(ori_grid[i], test.grid[i])
		}

		if res := orangesRotting(test.grid); res != test.maxTime {
			t.Errorf("orangeRotting(%v) = %d,want %d\n", ori_grid, res, test.maxTime)
		}
	}
}

func TestOrangeRotting2(t *testing.T) {
	tests := []struct {
		grid    [][]int
		maxTime int
	}{
		{[][]int{{}}, -1},
		{[][]int{{1}}, -1},
		{[][]int{{2}}, 0},
		{[][]int{{0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{1, 2}}, 1},
		{[][]int{{0}, {1}}, -1},
		{[][]int{{0}, {2}}, 0},
		{[][]int{{1}, {2}}, 1},
		{[][]int{{0, 0, 0}}, 0},
		{[][]int{{0, 0, 1}}, -1},
		{[][]int{{0, 0, 2}}, 0},
		{[][]int{{0, 1, 0}}, -1},
		{[][]int{{0, 1, 1}}, -1},
		{[][]int{{0, 1, 2}}, 1},
		{[][]int{{0, 2, 0}}, 0},
		{[][]int{{0, 2, 1}}, 1},
		{[][]int{{1, 0, 0}}, -1},
		{[][]int{{1, 0, 1}}, -1},
		{[][]int{{1, 0, 2}}, -1},
		{[][]int{{1, 1, 0}}, -1},
		{[][]int{{1, 1, 1}}, -1},
		{[][]int{{1, 1, 2}}, 2},
		{[][]int{{1, 2, 0}}, 1},
		{[][]int{{1, 2, 1}}, 1},
		{[][]int{{1, 2, 2}}, 1},
		{[][]int{{2, 0, 0}}, 0},
		{[][]int{{2, 0, 1}}, -1},
		{[][]int{{2, 0, 2}}, 0},
		{[][]int{{2, 1, 0}}, 1},
		{[][]int{{2, 1, 1}}, 2},
		{[][]int{{2, 1, 2}}, 1},
		{[][]int{{2, 2, 0}}, 0},
		{[][]int{{2, 2, 1}}, 1},
		{[][]int{{2, 2, 2}}, 0},

		{[][]int{{0}, {0}, {0}}, 0},
		{[][]int{{0}, {0}, {1}}, -1},
		{[][]int{{0}, {0}, {2}}, 0},
		{[][]int{{0}, {1}, {0}}, -1},
		{[][]int{{0}, {1}, {1}}, -1},
		{[][]int{{0}, {1}, {2}}, 1},
		{[][]int{{0}, {2}, {0}}, 0},
		{[][]int{{0}, {2}, {1}}, 1},
		{[][]int{{1}, {0}, {0}}, -1},
		{[][]int{{1}, {0}, {1}}, -1},
		{[][]int{{1}, {0}, {2}}, -1},
		{[][]int{{1}, {1}, {0}}, -1},
		{[][]int{{1}, {1}, {1}}, -1},
		{[][]int{{1}, {1}, {2}}, 2},
		{[][]int{{1}, {2}, {0}}, 1},
		{[][]int{{1}, {2}, {1}}, 1},
		{[][]int{{1}, {2}, {2}}, 1},
		{[][]int{{2}, {0}, {0}}, 0},
		{[][]int{{2}, {0}, {1}}, -1},
		{[][]int{{2}, {0}, {2}}, 0},
		{[][]int{{2}, {1}, {0}}, 1},
		{[][]int{{2}, {1}, {1}}, 2},
		{[][]int{{2}, {1}, {2}}, 1},
		{[][]int{{2}, {2}, {0}}, 0},
		{[][]int{{2}, {2}, {1}}, 1},
		{[][]int{{2}, {2}, {2}}, 0},

		{[][]int{{0, 0}, {0, 0}}, 0},
		{[][]int{{0, 0}, {0, 1}}, -1},
		{[][]int{{0, 0}, {0, 2}}, 0},
		{[][]int{{0, 0}, {1, 0}}, -1},
		{[][]int{{0, 0}, {1, 1}}, -1},
		{[][]int{{0, 0}, {1, 2}}, 1},
		{[][]int{{0, 0}, {2, 0}}, 0},
		{[][]int{{0, 0}, {2, 1}}, 1},
		{[][]int{{0, 0}, {2, 2}}, 0},
		{[][]int{{0, 1}, {0, 0}}, -1},
		{[][]int{{0, 1}, {0, 1}}, -1},
		{[][]int{{0, 1}, {1, 0}}, -1},
		{[][]int{{0, 1}, {1, 1}}, -1},
		{[][]int{{0, 1}, {1, 2}}, 1},
		{[][]int{{0, 1}, {2, 0}}, -1},
		{[][]int{{0, 1}, {2, 1}}, 2},
		{[][]int{{0, 1}, {2, 2}}, 1},
		{[][]int{{0, 2}, {0, 0}}, 0},
		{[][]int{{0, 2}, {0, 1}}, 1},
		{[][]int{{0, 2}, {0, 2}}, 0},
		{[][]int{{0, 2}, {1, 0}}, -1},
		{[][]int{{0, 2}, {1, 1}}, 2},
		{[][]int{{0, 2}, {1, 2}}, 1},
		{[][]int{{0, 2}, {2, 0}}, 0},
		{[][]int{{0, 2}, {2, 1}}, 1},
		{[][]int{{0, 2}, {2, 2}}, 0},
		{[][]int{{1, 0}, {0, 0}}, -1},
		{[][]int{{1, 0}, {0, 1}}, -1},
		{[][]int{{1, 0}, {0, 2}}, -1},
		{[][]int{{1, 0}, {1, 0}}, -1},
		{[][]int{{1, 0}, {1, 1}}, -1},
		{[][]int{{1, 0}, {1, 2}}, 2},
		{[][]int{{1, 0}, {2, 0}}, 1},
		{[][]int{{1, 0}, {2, 1}}, 1},
		{[][]int{{1, 0}, {2, 2}}, 1},
		{[][]int{{1, 1}, {0, 0}}, -1},
		{[][]int{{1, 1}, {0, 1}}, -1},
		{[][]int{{1, 1}, {0, 2}}, 2},
		{[][]int{{1, 1}, {1, 0}}, -1},
		{[][]int{{1, 1}, {1, 1}}, -1},
		{[][]int{{1, 1}, {1, 2}}, 2},
		{[][]int{{1, 1}, {2, 0}}, 2},
		{[][]int{{1, 1}, {2, 1}}, 2},
		{[][]int{{1, 1}, {2, 2}}, 1},
		{[][]int{{1, 2}, {0, 0}}, 1},
		{[][]int{{1, 2}, {0, 1}}, 1},
		{[][]int{{1, 2}, {0, 2}}, 1},
		{[][]int{{1, 2}, {1, 0}}, 2},
		{[][]int{{1, 2}, {1, 1}}, 2},
		{[][]int{{1, 2}, {1, 2}}, 1},
		{[][]int{{1, 2}, {2, 0}}, 1},
		{[][]int{{1, 2}, {2, 1}}, 1},
		{[][]int{{1, 2}, {2, 2}}, 1},
		{[][]int{{2, 0}, {0, 0}}, 0},
		{[][]int{{2, 0}, {0, 1}}, -1},
		{[][]int{{2, 0}, {0, 2}}, 0},
		{[][]int{{2, 0}, {1, 0}}, 1},
		{[][]int{{2, 0}, {1, 1}}, 2},
		{[][]int{{2, 0}, {1, 2}}, 1},
		{[][]int{{2, 0}, {2, 0}}, 0},
		{[][]int{{2, 0}, {2, 1}}, 1},
		{[][]int{{2, 0}, {2, 2}}, 0},
		{[][]int{{2, 1}, {0, 0}}, 1},
		{[][]int{{2, 1}, {0, 1}}, 2},
		{[][]int{{2, 1}, {0, 2}}, 1},
		{[][]int{{2, 1}, {1, 0}}, 1},
		{[][]int{{2, 1}, {1, 1}}, 2},
		{[][]int{{2, 1}, {1, 2}}, 1},
		{[][]int{{2, 1}, {2, 0}}, 1},
		{[][]int{{2, 1}, {2, 1}}, 1},
		{[][]int{{2, 1}, {2, 2}}, 1},
		{[][]int{{2, 2}, {0, 0}}, 0},
		{[][]int{{2, 2}, {0, 1}}, 1},
		{[][]int{{2, 2}, {0, 2}}, 0},
		{[][]int{{2, 2}, {1, 0}}, 1},
		{[][]int{{2, 2}, {1, 1}}, 1},
		{[][]int{{2, 2}, {1, 2}}, 1},
		{[][]int{{2, 2}, {2, 0}}, 0},
		{[][]int{{2, 2}, {2, 1}}, 1},
		{[][]int{{2, 2}, {2, 2}}, 0},

		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, -1},
		{[][]int{{2, 0, 0}, {0, 1, 1}, {0, 0, 0}, {0, 1, 1}}, -1},
		{[][]int{{2, 0, 1}, {0, 1, 1}, {0, 2, 0}, {0, 1, 1}}, 3},
	}
	for _, test := range tests {
		// printMatrix := func(matrix [][]int) {
		// 	for _, line := range matrix {
		// 		t.Logf("%v\n", line)
		// 	}
		// }
		var ori_grid [][]int = make([][]int, len(test.grid))
		for i := range ori_grid {
			ori_grid[i] = make([]int, len(test.grid[0]))
			copy(ori_grid[i], test.grid[i])
		}

		if res := orangesRotting2(test.grid); res != test.maxTime {
			t.Errorf("orangeRotting(%v) = %d,want %d\n", ori_grid, res, test.maxTime)
		}
	}
}
