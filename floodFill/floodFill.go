/*
 * File: \floodFill\floodFill.go                                               *
 * Project: leetcode                                                           *
 * Created At: Wednesday, 2021/09/29 , 23:44:55                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/03/15 , 13:20:35                               *
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

// floodFill 递归DFS方式，使用辅助函数adj获取单个节点的邻节点
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	l := len(image)
	if l == 0 {
		return image
	}
	c := len(image[0])
	if c == 0 {
		return image
	}

	ori_color := image[sr][sc]
	// adj 返回(sr,sc)点上下左右四个方向上的邻接点坐标
	adj := func(sr, sc int) [][]int {
		res := [][]int{}
		lmax := l - 1
		cmax := c - 1
		if l < 1 || c < 1 {
			return res
		}

		if lmax == 0 {
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

		if cmax == 0 {
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

		if sr == 0 {
			if sc == 0 {
				res = append(res, []int{0, 1}, []int{1, 0})
			} else if sc == cmax {
				res = append(res, []int{0, cmax - 1}, []int{1, cmax})
			} else {
				res = append(res, []int{0, sc - 1}, []int{0, sc + 1}, []int{1, sc})
			}
		} else if sr == lmax {
			if sc == 0 {
				res = append(res, []int{lmax, 1}, []int{lmax - 1, 0})
			} else if sc == cmax {
				res = append(res, []int{lmax, cmax - 1}, []int{lmax - 1, cmax})
			} else {
				res = append(res, []int{lmax, sc - 1}, []int{lmax, sc + 1}, []int{lmax - 1, sc})
			}
		} else {
			if sc == 0 {
				res = append(res, []int{sr, 1}, []int{sr - 1, 0}, []int{sr + 1, 0})
			} else if sc == cmax {
				res = append(res, []int{sr, cmax - 1}, []int{sr - 1, cmax}, []int{sr + 1, cmax})
			} else {
				res = append(res, []int{sr, sc - 1}, []int{sr, sc + 1}, []int{sr - 1, sc}, []int{sr + 1, sc})
			}
		}

		return res
	}

	var dfs func(sr int, sc int)
	dfs = func(sr int, sc int) {
		if image[sr][sc] == newColor {
			return
		}
		color := image[sr][sc]

		if color == ori_color {
			image[sr][sc] = newColor
		} else {
			return
		}

		for _, a := range adj(sr, sc) {
			if len(a) == 0 {
				return
			}
			dfs(a[0], a[1])
		}

	}
	dfs(sr, sc)

	return image
}

// floodFill2 迭代BFS方式
func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	rl := len(image)

	if rl == 0 || len(image[0]) == 0 || image[sr][sc] == newColor {
		return image
	}
	cl := len(image[0])
	color := image[sr][sc]
	image[sr][sc] = newColor

	nodes := [][]int{{sr, sc}}
	var (
		dx = []int{1, 0, 0, -1}
		dy = []int{0, 1, -1, 0}
	)

	for i := 0; i < len(nodes); i++ {
		cor := nodes[i]
		for j := 0; j < 4; j++ {
			x, y := cor[0]+dx[j], cor[1]+dy[j]
			if x > -1 && y > -1 && x < rl && y < cl && image[x][y] == color {
				nodes = append(nodes, []int{x, y})
				image[x][y] = newColor
			}
		}

	}

	return image
}

// floodFill3 递归 DFS方式
func floodFill3(image [][]int, sr int, sc int, newColor int) [][]int {
	l := len(image)

	if l == 0 || len(image[0]) == 0 {
		return image
	}

	color := image[sr][sc]
	if color == newColor {
		return image
	}

	lmax := len(image) - 1
	cmax := len(image[0]) - 1

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i > lmax || j > cmax || image[i][j] != color {
			return
		}
		image[i][j] = newColor
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	dfs(sr, sc)

	return image
}

func main() {

	fmt.Println("--------------------floodFill begin---------------------")
	funcs := map[string]func(image [][]int, sr int, sc int, newColor int) [][]int{
		"floodFill1": floodFill1,
		"floodFill2": floodFill2,
		"floodFill3": floodFill3,
	}
	tests2 := []struct {
		image, res       [][]int
		sr, sc, newColor int
	}{
		{[][]int{}, [][]int{}, 0, 0, 2},
		{[][]int{}, [][]int{}, 1, 2, 2},
		{[][]int{{1}}, [][]int{{2}}, 0, 0, 2},
		{[][]int{{0, 1}}, [][]int{{0, 2}}, 0, 1, 2},
		{[][]int{{0}, {1}}, [][]int{{0}, {2}}, 1, 0, 2},
		{[][]int{{0, 1, 1}}, [][]int{{2, 1, 1}}, 0, 0, 2},
		{[][]int{{0, 1, 1}}, [][]int{{0, 2, 2}}, 0, 1, 2},
		{[][]int{{0, 1, 1}}, [][]int{{1, 2, 2}}, 0, 2, 2},
		{[][]int{{1}, {0}, {0}}, [][]int{{2}, {0}, {0}}, 0, 0, 2},
		{[][]int{{1}, {0}, {0}}, [][]int{{1}, {2}, {2}}, 1, 0, 2},
		{[][]int{{1}, {0}, {0}}, [][]int{{1}, {2}, {2}}, 2, 0, 2},
		{[][]int{{1, 0}, {0, 1}}, [][]int{{2, 0}, {0, 1}}, 0, 0, 2},
		{[][]int{{1, 0}, {0, 1}}, [][]int{{1, 2}, {0, 1}}, 0, 1, 2},
		{[][]int{{1, 0}, {0, 1}}, [][]int{{1, 0}, {2, 1}}, 1, 0, 2},
		{[][]int{{1, 0}, {0, 1}}, [][]int{{1, 0}, {0, 2}}, 1, 1, 2},
		{[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 0, 0, 2},
		{[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 1, 0, 2},
		{[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 0, 1, 2},
		{[][]int{{0, 0}, {0, 0}}, [][]int{{2, 2}, {2, 2}}, 1, 1, 2},
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, 1, 1, 2},
		{[][]int{{0, 0, 0}, {0, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1},
	}
	for fn, f := range funcs {
		fmt.Printf("--------------------%s begin---------------------\n", fn)
		for _, test := range tests2 {
			fmt.Println("before:")
			printMatrix(test.image)
			image := f(test.image, test.sr, test.sc, test.newColor)
			fmt.Println("after:")
			printMatrix(image)
			if !reflect.DeepEqual(image, test.res) {
				fmt.Println("failed: want")
				printMatrix(test.res)
			} else {
				fmt.Println("passed")
			}
		}

		fmt.Printf("--------------------%s end---------------------\n", fn)
	}
	fmt.Println("--------------------floodFill end---------------------")
}

func printMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Printf("%v\n", line)
	}
}
