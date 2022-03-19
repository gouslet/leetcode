/*
 * File: \permute\permute.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Friday, 2022/03/18 , 17:21:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 17:36:02                              *
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

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	l := len(nums)
	if l == 1 {
		res = append(res, nums)
	}

	for i, n := range nums {
		sub := make([]int, l)
		copy(sub, nums)
		sub[i], sub[0] = sub[0], sub[i]

		subres := permute(sub[1:])
		for j := range subres {
			subres[j] = append(subres[j], n)
		}
		res = append(res, subres...)
	}
	return res
}

func main() {
	tests := []struct {
		nums []int
		res  [][]int
	}{
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {3, 1, 2}}},
	}

	for _, test := range tests {
		res := permute(test.nums)
		if !reflect.DeepEqual(res, test.res) {
			fmt.Printf("permute(%d) = %v,want %v\n", test.nums, res, test.res)
		}
	}
}
