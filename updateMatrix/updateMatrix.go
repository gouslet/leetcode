/*
 * File: \updateMatrix\updateMatrix.go                                         *
 * Project: leetcode                                                           *
 * Created At: Friday, 2021/10/15 , 23:07:39                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/03/15 , 20:30:44                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"reflect"
)

func updateMatrix(mat [][]int) [][]int {
	r := len(mat)
	if r == 0 {
		return nil
	}
	c := len(mat[0])
	if c == 0 {
		return nil
	}
	nodes := [][]int{}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			if mat[i][j] == 0 {
				nodes = append(nodes, []int{i, j})
			} else {
				res[i][j] = -1
			}
		}
	}

	for i := 0; i < len(nodes); i++ {
		cell := nodes[i]
		a := cell[0]
		b := cell[1]

		if a > 0 && res[a-1][b] < 0 {
			res[a-1][b] = res[a][b] + 1
			nodes = append(nodes, []int{a - 1, b})
		}

		if a+1 < r && res[a+1][b] < 0 {
			res[a+1][b] = res[a][b] + 1
			nodes = append(nodes, []int{a + 1, b})

		}

		if b+1 < c && res[a][b+1] < 0 {
			res[a][b+1] = res[a][b] + 1
			nodes = append(nodes, []int{a, b + 1})
		}

		if b > 0 && res[a][b-1] < 0 {
			res[a][b-1] = res[a][b] + 1
			nodes = append(nodes, []int{a, b - 1})
		}
	}

	return res
}

// func updateMatrix(mat [][]int) [][]int {
// 	r := len(mat)
// 	if r == 0 {
// 		return nil
// 	}
// 	c := len(mat[0])
// 	if c == 0 {
// 		return nil
// 	}

// 	for i := 0; i < r; i++ {
// 		for j := 0; j < c; j++ {
// 			if mat[i][j] != 0 {
// 				visited := [][]bool{}
// 				for i := 0; i < r; i++ {
// 					visited = append(visited, make([]bool, c))
// 				}
// 				mat[i][j] = dfs(mat, i, j, visited)
// 			}
// 		}
// 	}

// 	return mat
// }

func dfs(mat [][]int, sr int, sc int, visited [][]bool) int {
	cells := [][]int{{sr, sc}}
	for len(cells) != 0 {
		v := cells[0]
		cells = cells[1:]
		adjs := adj(mat, v[0], v[1])
		for _, a := range adjs {
			if len(a) == 0 {
				return 0
			}
			if !visited[a[0]][a[1]] {
				visited[a[0]][a[1]] = true
				cells = append(cells, []int{a[0], a[1]})
				if mat[a[0]][a[1]] == 0 {
					if a[0] > sr {
						if a[1] > sc {
							return a[0] - sr + a[1] - sc
						} else {
							return a[0] - sr + sc - a[1]
						}
					} else {
						if a[1] > sc {
							return sr - a[0] + a[1] - sc
						} else {
							return sr - a[0] + sc - a[1]
						}
					}
				}
			}
		}
	}
	return -1
}

// adj 返回mat矩阵中点(sr,sc)的相邻点的坐标
func adj(mat [][]int, sr, sc int) [][]int {
	res := [][]int{}
	lmax := len(mat) - 1
	if lmax < 0 {
		return res
	}
	cmax := len(mat[0]) - 1
	if cmax < 0 {
		return res
	}

	if lmax == 0 { // 1 × n 矩阵
		if cmax == 0 {
			return res
		} else {
			if sc == 0 {
				return [][]int{{0, 1}}
			} else if sc == cmax {
				return [][]int{{0, cmax - 1}}
			} else {
				return [][]int{{0, sc - 1}, {0, sc + 1}}
			}
		}
	}

	if cmax == 0 { // m × 1 矩阵
		if lmax == 0 {
			return res
		} else {
			if sr == 0 {
				return [][]int{{1, 0}}
			} else if sr == lmax {
				return [][]int{{lmax - 1, 0}}
			} else {
				return [][]int{{sr - 1, 0}, {sr + 1, 0}}
			}
		}
	}

	// m × n 矩阵
	if sr == 0 {
		if sc == 0 { // 点(0,0)
			res = append(res, []int{0, 1}, []int{1, 0})
		} else if sc == cmax { // 点(0,cmax)
			res = append(res, []int{0, cmax - 1}, []int{1, cmax})
		} else { // 点(0,?)
			res = append(res, []int{0, sc - 1}, []int{0, sc + 1}, []int{1, sc})
		}
	} else if sr == lmax {
		if sc == 0 { // 点(0,lmax)
			res = append(res, []int{lmax, 1}, []int{lmax - 1, 0})
		} else if sc == cmax { // 点(lmax,cmax)
			res = append(res, []int{lmax, cmax - 1}, []int{lmax - 1, cmax})
		} else { // 点(lmax,?)
			res = append(res, []int{lmax, sc - 1}, []int{lmax, sc + 1}, []int{lmax - 1, sc})
		}
	} else {
		if sc == 0 { // 点(?,0)
			res = append(res, []int{sr, 1}, []int{sr - 1, 0}, []int{sr + 1, 0})
		} else if sc == cmax { // 点(?,cmax)
			res = append(res, []int{sr, cmax - 1}, []int{sr - 1, cmax}, []int{sr + 1, cmax})
		} else { // 点(?,?)
			res = append(res, []int{sr, sc - 1}, []int{sr, sc + 1}, []int{sr - 1, sc}, []int{sr + 1, sc})
		}
	}

	return res
}

// printMatrix 输出矩阵
func printMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Printf("%v\n", line)
	}
}

func main() {
	fmt.Println("--------------------updateMatrix begin---------------------")
	tests := []struct {
		grid, res [][]int
	}{
		{[][]int{}, [][]int{}},
		{[][]int{{1, 0}}, [][]int{{1, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{0, 0, 0}},
			[][]int{
				{16, 16, 16},
				{15, 15, 15},
				{14, 14, 14},
				{13, 13, 13},
				{12, 12, 12},
				{11, 11, 11},
				{10, 10, 10},
				{9, 9, 9},
				{8, 8, 8},
				{7, 7, 7},
				{6, 6, 6},
				{5, 5, 5},
				{4, 4, 4},
				{3, 3, 3},
				{2, 2, 2},
				{1, 1, 1},
				{0, 0, 0},
			},
		},

		{
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			[][]int{
				{3, 3, 3},
				{2, 2, 2},
				{1, 1, 1},
				{0, 0, 0},
			},
		},
		{
			[][]int{{1}},
			[][]int{{0}},
		},
		// {[][]int{{0, 1}}, 0, 1, 2},
		// {[][]int{{0}, {1}}, [][]int{{0}, {2}}, 1, 0, 2},
		// {[][]int{{0, 1, 1}}, [][]int{{2, 1, 1}}, 0, 0, 2},
		// {[][]int{{0, 1, 1}}, [][]int{{0, 2, 2}}, 0, 1, 2},
		// {[][]int{{0, 1, 1}}, [][]int{{1, 2, 2}}, 0, 2, 2},
		// {[][]int{{1}, {0}, {0}}, [][]int{{2}, {0}, {0}}, 0, 0, 2},
		// {[][]int{{1}, {0}, {0}}, [][]int{{1}, {2}, {2}}, 1, 0, 2},
		// {[][]int{{1}, {0}, {0}}, [][]int{{1}, {2}, {2}}, 2, 0, 2},
		// {[][]int{{1, 0}, {0, 1}}, [][]int{{2, 0}, {0, 1}}, 0, 0, 2},
		// {[][]int{{1, 0}, {0, 1}}, [][]int{{1, 2}, {0, 1}}, 0, 1, 2},
		// {[][]int{{1, 0}, {0, 1}}, [][]int{{1, 0}, {2, 1}}, 1, 0, 2},
		// {[][]int{{1, 0}, {0, 1}}, [][]int{{1, 0}, {0, 2}}, 1, 1, 2},
		// {[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 0, 0, 2},
		// {[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 1, 0, 2},
		// {[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 0, 1, 2},
		// {[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 1, 1, 2},
		// {[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, 1, 1, 2},
		// {[][]int{{0, 0, 0}, {0, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1},
		// {[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
	}
	for _, test := range tests {
		shortestPaths := updateMatrix(test.grid)
		fmt.Println("originalMatrix: ")
		printMatrix(test.grid)
		fmt.Println("updatedMatrix: ")

		printMatrix(shortestPaths)
		if !reflect.DeepEqual(test.res, shortestPaths) {
			fmt.Println("failed: want")
			printMatrix(test.res)
		} else {
			fmt.Println("passed")
		}
	}
	fmt.Println("--------------------maxAreaOfIsland end---------------------")
}
