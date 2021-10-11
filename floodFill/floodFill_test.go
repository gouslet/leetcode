package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdj(t *testing.T) {
	tests := []struct {
		image, res [][]int
		sr, sc     int
	}{
		{[][]int{}, [][]int{}, 0, 0},
		{[][]int{}, [][]int{}, 1, 2},
		{[][]int{{1}}, [][]int{}, 0, 0},
		{[][]int{{1, 2, 3}}, [][]int{{0, 0}, {0, 2}}, 0, 1},
		{[][]int{{1}, {2}, {3}}, [][]int{{0, 0}, {2, 0}}, 1, 0},
	}

	for _, test := range tests {
		adjs := adj(test.image, test.sr, test.sc)
		if !reflect.DeepEqual(adjs, test.res) {
			t.Errorf("adj failed: %v got,%v wanted\n", adjs, test.res)
			break
		}
	}

}

func TestFloodFill(t *testing.T) {
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
	for _, test := range tests {
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
}
