package main

import (
	"fmt"
	"reflect"
)

func maxAreaOfIsland(grid [][]int) int {
	l := len(grid)
	if l == 0 {
		return 0
	}
	c := len(grid[0])
	if c == 0 {
		return 0
	}

	visited := [][]bool{}
	for i := 0; i < l; i++ {
		visited = append(visited, make([]bool, c))
	}

	area := 0
	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			if !visited[i][j] {
				visited[i][j] = true
				if (j == 0 || grid[i][j-1] == 0) && grid[i][j] == 1 {
					if newArea := dfs(grid, i, j, visited); newArea > area {
						area = newArea
					}
				}
			}
		}
	}

	return area
}

func dfs(grid [][]int, sr int, sc int, flag [][]bool) int {
	area := 0
	for _, a := range adj(grid, sr, sc) {
		if len(a) == 0 {
			return 0
		}
		if !flag[a[0]][a[1]] && grid[a[0]][a[1]] == 1 {
			flag[a[0]][a[1]] = true
			area += dfs(grid, a[0], a[1], flag)
		}
	}
	flag[sr][sc] = true
	area++

	return area
}

func adj(grid [][]int, sr, sc int) [][]int {
	res := [][]int{}
	lmax := len(grid) - 1
	if lmax < 0 {
		return res
	}
	cmax := len(grid[0]) - 1
	if cmax < 0 {
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

func main() {
	fmt.Println("--------------------adj begin---------------------")
	tests1 := []struct {
		grid, res [][]int
		sr, sc    int
	}{
		{[][]int{}, [][]int{}, 0, 0},
		{[][]int{}, [][]int{}, 1, 2},
		{[][]int{{1}}, [][]int{}, 0, 0},
		{[][]int{{1, 2, 3}}, [][]int{{0, 0}, {0, 2}}, 0, 1},
		{[][]int{{1}, {2}, {3}}, [][]int{{0, 0}, {2, 0}}, 1, 0},
	}

	for _, test := range tests1 {
		adjs := adj(test.grid, test.sr, test.sc)
		fmt.Printf("adj(%v,%d,%d) = %v: ", test.grid, test.sr, test.sc, adjs)
		if !reflect.DeepEqual(adjs, test.res) {
			fmt.Printf("failed: want %v\n", test.res)
		} else {
			fmt.Println("passed")
		}
	}
	fmt.Println("--------------------adj end---------------------")

	fmt.Println("--------------------maxAreaOfIsland begin---------------------")
	tests2 := []struct {
		grid [][]int
		res  int
	}{
		// {[][]int{}, 0, 0, 2},
		// {[][]int{}, 1, 2, 2},
		// {[][]int{{1}}, 0, 0, 2},
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
		{[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
	}
	for _, test := range tests2 {
		maxArea := maxAreaOfIsland(test.grid)
		fmt.Println("maxAreaOfIsland: ")
		printMatrix(test.grid)
		fmt.Println("maxAreaOfIsland =", maxArea)
		if maxArea != test.res {
			fmt.Printf("failed: want %d\n", test.res)
		} else {
			fmt.Println("passed")
		}
	}
	fmt.Println("--------------------maxAreaOfIsland end---------------------")
}

func printMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Printf("%v\n", line)
	}
}
