/*
 * File: \orangesRotting\orangeRotting.go                                      *
 * Project: leetcode                                                           *
 * Created At: Wednesday, 2022/03/16 , 10:48:45                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/03/17 , 01:14:45                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

func orangesRotting(grid [][]int) int {
	r := len(grid)
	if grid == nil || r == 0 {
		return -1
	}

	c := len(grid[0])
	if c == 0 {
		return -1
	}

	rotten := make([][]int, 0)
	k, fresh := 0, 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				fresh++
			} else if cell == 2 {
				rotten = append(rotten, []int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	for i := 0; i < len(rotten); i++ {
		cell := rotten[i]
		a, b := cell[0], cell[1]

		if a > 0 && grid[a-1][b] == 1 {
			grid[a-1][b] = grid[a][b] + 1
			rotten = append(rotten, []int{a - 1, b})
			fresh--
			if grid[a-1][b] > k {
				k = grid[a-1][b]
			}
		}

		if a < r-1 && grid[a+1][b] == 1 {
			grid[a+1][b] = grid[a][b] + 1
			fresh--
			rotten = append(rotten, []int{a + 1, b})
			if grid[a+1][b] > k {
				k = grid[a+1][b]
			}
		}

		if b > 0 && grid[a][b-1] == 1 {
			grid[a][b-1] = grid[a][b] + 1
			fresh--
			rotten = append(rotten, []int{a, b - 1})
			if grid[a][b-1] > k {
				k = grid[a][b-1]
			}
		}

		if b < c-1 && grid[a][b+1] == 1 {
			grid[a][b+1] = grid[a][b] + 1
			fresh--
			rotten = append(rotten, []int{a, b + 1})
			if grid[a][b+1] > k {
				k = grid[a][b+1]
			}
		}
	}

	if fresh == 0 {
		return k - 2
	} else {
		return -1
	}
}

func orangesRotting2(grid [][]int) int {
	r := len(grid)
	if grid == nil || r == 0 {
		return -1
	}

	c := len(grid[0])
	if c == 0 {
		return -1
	}

	rotten := make([][]int, 0)
	k, fresh := 0, 0

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				fresh++
			} else if cell == 2 {
				rotten = append(rotten, []int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	for len(rotten) > 0 {
		for i, l := 0, len(rotten); i < l; i++ {
			cell := rotten[0]
			rotten = rotten[1:]
			a, b := cell[0], cell[1]

			if a > 0 && grid[a-1][b] == 1 {
				grid[a-1][b] = 2
				rotten = append(rotten, []int{a - 1, b})
				fresh--
			}

			if a < r-1 && grid[a+1][b] == 1 {
				grid[a+1][b] = grid[a][b] + 1
				fresh--
				rotten = append(rotten, []int{a + 1, b})
			}

			if b > 0 && grid[a][b-1] == 1 {
				grid[a][b-1] = grid[a][b] + 1
				fresh--
				rotten = append(rotten, []int{a, b - 1})
			}

			if b < c-1 && grid[a][b+1] == 1 {
				grid[a][b+1] = grid[a][b] + 1
				fresh--
				rotten = append(rotten, []int{a, b + 1})
			}
		}

		if len(rotten) > 0 {
			k++
		}

	}

	if fresh == 0 {
		return k
	} else {
		return -1
	}
}
