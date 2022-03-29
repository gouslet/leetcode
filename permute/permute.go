/*
 * File: \permute\permute.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Friday, 2022/03/18 , 17:21:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/27 , 00:27:48                                *
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

func permute1(nums []int) [][]int {
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

		subres := permute1(sub[1:])
		for j := range subres {
			subres[j] = append(subres[j], n)
		}
		res = append(res, subres...)
	}
	return res
}

// permute2 回溯法
func permute2(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	l := len(nums)

	var permutation func(cur int)
	permutation = func(cur int) {
		if cur == l {
			output := make([]int, l)
			copy(output, nums)
			res = append(res, output)
			return
		}

		permutation(cur + 1)
		for i := cur + 1; i < l; i++ {
			nums[i], nums[cur] = nums[cur], nums[i]
			permutation(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}

	permutation(0)

	return res
}

// permute3 回溯法 优化
func permute3(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	l := len(nums)

	var fact func(int) int
	fact = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}

	res := make([][]int, 0, fact(l))

	var permutation func(cur int)
	permutation = func(cur int) {
		if cur == l {
			res = append(res, append([]int(nil), nums...))
			return
		}

		for i := cur; i < l; i++ {
			nums[i], nums[cur] = nums[cur], nums[i]
			permutation(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}

	permutation(0)

	return res
}

func main() {
	funcs := map[string]func(nums []int) [][]int{
		"permute1": permute1,
		"permute2": permute2,
		"permute3": permute3,
	}

	tests := []struct {
		nums []int
		res  [][]int
	}{
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}

	for n, f := range funcs {
		for _, test := range tests {
			res := f(test.nums)
			if !reflect.DeepEqual(res, test.res) {
				fmt.Printf("%s(%d) = %v,want %v\n", n, test.nums, res, test.res)
			}
		}
	}
}
