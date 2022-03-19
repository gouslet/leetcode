/*
 * File: \permute\permute.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Friday, 2022/03/18 , 17:21:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/03/18 , 17:54:57                                *
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

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		res  [][]int
	}{
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {3, 1, 2}}},
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
		res := permute(test.nums)
		sort2DInts(res)
		sort2DInts(test.res)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("permute(%v) = %v,want %v", test.nums, res, test.res)
		}
	}
}
