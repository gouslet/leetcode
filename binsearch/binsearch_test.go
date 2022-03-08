/*
 * File: \binsearch\binsearch_test.go                                          *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/7 , 13:19:50                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 13:25:56                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"testing"
)

func TestSearch1(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    int
	}{
		{nil, 55, -1},
		{[]int{}, 5, -1},
		{[]int{0}, -5, -1},
		{[]int{0}, 0, 0},
		{[]int{0}, 5, -1},
		{[]int{-3, -2, -1}, -5, -1},
		{[]int{-3, -2, -1}, -2, 1},
		{[]int{-3, -2, -1}, 5, -1},
		{[]int{1, 2, 3}, -5, -1},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 5, -1},
		{[]int{-9999, 1, 2, 3, 9999}, -10000, -1},
		{[]int{-9999, 1, 2, 3, 9999}, -9999, 0},
		{[]int{-9999, 1, 2, 3, 9999}, 9999, 4},
		{[]int{-9999, 1, 2, 3, 9999}, 10000, -1},
	}

	for _, test := range tests {
		res := search1(test.nums, test.target)

		if res != test.res {
			t.Errorf("search1(%v,%d) = %d: want %d", test.nums, test.target, res, test.res)
		}
	}
}

func TestSearch2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    int
	}{
		{nil, 55, -1},
		{[]int{}, 5, -1},
		{[]int{0}, -5, -1},
		{[]int{0}, 0, 0},
		{[]int{0}, 5, -1},
		{[]int{-3, -2, -1}, -5, -1},
		{[]int{-3, -2, -1}, -2, 1},
		{[]int{-3, -2, -1}, 5, -1},
		{[]int{1, 2, 3}, -5, -1},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 5, -1},
		{[]int{-9999, 1, 2, 3, 9999}, -10000, -1},
		{[]int{-9999, 1, 2, 3, 9999}, -9999, 0},
		{[]int{-9999, 1, 2, 3, 9999}, 9999, 4},
		{[]int{-9999, 1, 2, 3, 9999}, 10000, -1},
	}

	for _, test := range tests {
		res := search2(test.nums, test.target)

		if res != test.res {
			t.Errorf("search2(%v,%d) = %d: want %d", test.nums, test.target, res, test.res)
		}
	}
}

func TestSearch3(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    int
	}{
		{nil, 55, -1},
		{[]int{}, 5, -1},
		{[]int{0}, -5, -1},
		{[]int{0}, 0, 0},
		{[]int{0}, 5, -1},
		{[]int{-3, -2, -1}, -5, -1},
		{[]int{-3, -2, -1}, -2, 1},
		{[]int{-3, -2, -1}, 5, -1},
		{[]int{1, 2, 3}, -5, -1},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 5, -1},
		{[]int{-9999, 1, 2, 3, 9999}, -10000, -1},
		{[]int{-9999, 1, 2, 3, 9999}, -9999, 0},
		{[]int{-9999, 1, 2, 3, 9999}, 9999, 4},
		{[]int{-9999, 1, 2, 3, 9999}, 10000, -1},
	}

	for _, test := range tests {
		res := search3(test.nums, test.target)

		if res != test.res {
			t.Errorf("search3(%v,%d) = %d: want %d", test.nums, test.target, res, test.res)
		}
	}
}

func TestSearch4(t *testing.T) {
	tests := []struct {
		nums   []int16
		target int16
		res    int16
	}{
		{nil, 55, -1},
		{[]int16{}, 5, -1},
		{[]int16{0}, -5, -1},
		{[]int16{0}, 0, 0},
		{[]int16{0}, 5, -1},
		{[]int16{-3, -2, -1}, -5, -1},
		{[]int16{-3, -2, -1}, -2, 1},
		{[]int16{-3, -2, -1}, 5, -1},
		{[]int16{1, 2, 3}, -5, -1},
		{[]int16{1, 2, 3}, 2, 1},
		{[]int16{1, 2, 3}, 5, -1},
		{[]int16{-9999, 1, 2, 3, 9999}, -10000, -1},
		{[]int16{-9999, 1, 2, 3, 9999}, -9999, 0},
		{[]int16{-9999, 1, 2, 3, 9999}, 9999, 4},
		{[]int16{-9999, 1, 2, 3, 9999}, 10000, -1},
	}

	for _, test := range tests {
		res := search4(test.nums, test.target)

		if res != test.res {
			t.Errorf("search4(%v,%d) = %d: want %d", test.nums, test.target, res, test.res)
		}
	}
}
