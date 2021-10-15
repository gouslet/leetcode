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

	// res := [][]int{}

	// for k := 0; k < r; k++ {
	// 	res = append(res, make([]int, c))
	// }

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] != 0 {
				// visited := [][]bool{}
				// for i := 0; i < r; i++ {
				// 	visited = append(visited, make([]bool, c))
				// }
				// mat[i][j] = dfs(mat, i, j, visited)
				mat[i][j] = dfs(mat, i, j)
			}
		}
	}

	return mat
}

// func dfs(mat [][]int, sr int, sc int, visited [][]bool) int {
// 	cells := [][]int{{sr, sc}}
// 	for len(cells) != 0 {
// 		v := cells[0]
// 		cells = cells[1:]
// 		adjs := adj(mat, v[0], v[1])
// 		for _, a := range adjs {
// 			if len(a) == 0 {
// 				return 0
// 			}
// 			if !visited[a[0]][a[1]] {
// 				visited[a[0]][a[1]] = true
// 				cells = append(cells, []int{a[0], a[1]})
// 				if mat[a[0]][a[1]] == 0 {
// 					if a[0] > sr {
// 						if a[1] > sc {
// 							return a[0] - sr + a[1] - sc
// 						} else {
// 							return a[0] - sr + sc - a[1]
// 						}
// 					} else {
// 						if a[1] > sc {
// 							return sr - a[0] + a[1] - sc
// 						} else {
// 							return sr - a[0] + sc - a[1]
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

func dfs(mat [][]int, sr int, sc int) int {
	cells := [][]int{{sr, sc}}
	for len(cells) != 0 {
		v := cells[0]
		cells = cells[1:]
		for _, a := range adj(mat, v[0], v[1]) {
			if len(a) == 0 {
				return 0
			}
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
	return -1
}

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
		// {[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
	}
	for _, test := range tests {
		shortestPaths := updateMatrix(test.grid)
		fmt.Println("originalMatrix: ")
		printMatrix(test.grid)
		fmt.Println("updatedMatrix: ")

		printMatrix(shortestPaths)
		if !reflect.DeepEqual(test.grid, shortestPaths) {
			fmt.Println("failed: want")
			printMatrix(test.res)
		} else {
			fmt.Println("passed")
		}
	}
	fmt.Println("--------------------maxAreaOfIsland end---------------------")
}
