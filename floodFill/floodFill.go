package main

import (
	"fmt"
	"reflect"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	l := len(image)
	if l == 0 {
		return image
	}
	c := len(image[0])
	if c == 0 {
		return image
	}

	color := image[sr][sc]
	flag := [][]bool{}
	for i := 0; i < l; i++ {
		flag = append(flag, make([]bool, c))
	}
	image = dfs(image, sr, sc, flag, color, newColor)

	return image
}

func dfs(image [][]int, sr int, sc int, flag [][]bool, color, newColor int) [][]int {
	for _, a := range adj(image, sr, sc) {
		if len(a) == 0 {
			return image
		}
		if !flag[a[0]][a[1]] && image[a[0]][a[1]] == color {
			image[a[0]][a[1]] = newColor
			flag[a[0]][a[1]] = true
			image = dfs(image, a[0], a[1], flag, color, newColor)
		}
	}
	flag[sr][sc] = true
	image[sr][sc] = newColor

	return image
}

func adj(image [][]int, sr, sc int) [][]int {
	res := [][]int{}
	lmax := len(image) - 1
	if lmax < 0 {
		return res
	}
	cmax := len(image[0]) - 1
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
		image, res [][]int
		sr, sc     int
	}{
		{[][]int{}, [][]int{}, 0, 0},
		{[][]int{}, [][]int{}, 1, 2},
		{[][]int{{1}}, [][]int{}, 0, 0},
		{[][]int{{1, 2, 3}}, [][]int{{0, 0}, {0, 2}}, 0, 1},
		{[][]int{{1}, {2}, {3}}, [][]int{{0, 0}, {2, 0}}, 1, 0},
	}

	for _, test := range tests1 {
		adjs := adj(test.image, test.sr, test.sc)
		fmt.Printf("adj(%v,%d,%d) = %v: ", test.image, test.sr, test.sc, adjs)
		if !reflect.DeepEqual(adjs, test.res) {
			fmt.Printf("failed: want %v\n", test.res)
		} else {
			fmt.Println("passed")
		}
	}
	fmt.Println("--------------------adj end---------------------")
	fmt.Println("--------------------floodFill begin---------------------")
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
	for _, test := range tests2 {
		fmt.Println("before:")
		printMatrix(test.image)
		image := floodFill(test.image, test.sr, test.sc, test.newColor)
		fmt.Println("after:")
		printMatrix(image)
		if !reflect.DeepEqual(image, test.res) {
			fmt.Println("failed: want")
			printMatrix(test.res)
		} else {
			fmt.Println("passed")
		}
	}
	fmt.Println("--------------------floodFill end---------------------")
}

func printMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Printf("%v\n", line)
	}
}
