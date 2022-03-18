/*
 * File: \combine\combine_test.go                                              *
 * Project: leetcode                                                           *
 * Created At: Thursday, 2022/03/17 , 18:12:54                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/03/18 , 13:16:12                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n, k int
		res  [][]int
	}{
		{4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
	}
	sort2DInts := func(nums [][]int) {
		for _, cell := range nums {
			sort.Ints(cell)
		}

		sort.Slice(
			nums,
			func(i, j int) bool {
				a := nums[i]
				b := nums[j]
				var l int
				var res bool
				if la, lb := len(a), len(b); la < lb {
					l = la
					res = true
				} else {
					l = lb
				}
				for x := 0; x < l; x++ {
					if a[x] < b[x] {
						return true
					}
				}
				return res
			},
		)

	}
	for _, test := range tests {
		res := combine(test.n, test.k)
		sort2DInts(res)
		sort2DInts(test.res)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("combine(%d,%d) = %v,want %v", test.n, test.k, res, test.res)
		}
	}
}

func TestCombine2(t *testing.T) {
	tests := []struct {
		n, k int
		res  [][]int
	}{
		{4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
	}
	sort2DInts := func(nums [][]int) {
		for _, cell := range nums {
			sort.Ints(cell)
		}

		sort.Slice(
			nums,
			func(i, j int) bool {
				a := nums[i]
				b := nums[j]
				var l int
				var res bool
				if la, lb := len(a), len(b); la < lb {
					l = la
					res = true
				} else {
					l = lb
				}
				for x := 0; x < l; x++ {
					if a[x] < b[x] {
						return true
					}
				}
				return res
			},
		)

	}
	for _, test := range tests {
		res := combine2(test.n, test.k)
		sort2DInts(res)
		sort2DInts(test.res)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("combine2(%d,%d) = %v,want %v", test.n, test.k, res, test.res)
		}
	}
}
