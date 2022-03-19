/*
 * File: \floodFill\floodFill_test.go                                          *
 * Project: leetcode                                                           *
 * Created At: Monday, 2021/10/11 , 18:23:48                                   *
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

import (
	"reflect"
	"testing"
)

// func TestAdj (t *testing.T) {
// 	tests := []struct {
// 		image, res [][]int
// 		sr, sc     int
// 	}{
// 		{[][]int{}, [][]int{}, 0, 0},
// 		{[][]int{}, [][]int{}, 1, 2},
// 		{[][]int{{1}}, [][]int{}, 0, 0},
// 		{[][]int{{1, 2, 3}}, [][]int{{0, 0}, {0, 2}}, 0, 1},
// 		{[][]int{{1}, {2}, {3}}, [][]int{{0, 0}, {2, 0}}, 1, 0},
// 	}

// 	for _, test := range tests {
// 		adjs := floodFill1.adj(test.sr, test.sc)
// 		if !reflect.DeepEqual(adjs, test.res) {
// 			t.Errorf("adj failed: %v got,%v wanted\n", adjs, test.res)
// 			break
// 		}
// 	}

// }

func TestFloodFill1(t *testing.T) {
	tests := []struct {
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
		{[][]int{{0, 1, 1}}, [][]int{{0, 2, 2}}, 0, 2, 2},
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

	for _, test := range tests {
		printMatrix := func(matrix [][]int) {
			for _, line := range matrix {
				t.Logf("%v\n", line)
			}
		}
		var ori_image [][]int = make([][]int, len(test.image))
		for i := range ori_image {
			ori_image[i] = make([]int, len(test.image[0]))
			copy(ori_image[i], test.image[i])
		}

		image := floodFill1(test.image, test.sr, test.sc, test.newColor)
		if !reflect.DeepEqual(image, test.res) {
			t.Log("before:")
			printMatrix(ori_image)
			t.Log("after:")
			printMatrix(test.image)
			t.Error("failed: want")
			printMatrix(test.res)
		}
	}
}

func TestFloodFill2(t *testing.T) {
	tests := []struct {
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
		{[][]int{{0, 1, 1}}, [][]int{{0, 2, 2}}, 0, 2, 2},
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
	for _, test := range tests {
		t.Log("before:")
		printMatrix(test.image)
		image := floodFill2(test.image, test.sr, test.sc, test.newColor)
		t.Log("after:")
		printMatrix(image)
		if !reflect.DeepEqual(image, test.res) {
			t.Error("failed: want")
			printMatrix(test.res)
		}
	}
}

func TestFloodFill3(t *testing.T) {
	tests := []struct {
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
		{[][]int{{0, 1, 1}}, [][]int{{0, 2, 2}}, 0, 2, 2},
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
	for _, test := range tests {
		t.Log("before:")
		printMatrix(test.image)
		image := floodFill3(test.image, test.sr, test.sc, test.newColor)
		t.Log("after:")
		printMatrix(image)
		if !reflect.DeepEqual(image, test.res) {
			t.Error("failed: want")
			printMatrix(test.res)
		}
	}
}
