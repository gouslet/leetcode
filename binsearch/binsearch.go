/*
 * File: \binsearch\binsearch.go                                               *
 * Project: leetcode                                                           *
 * Created At: Thursday, 2021/10/7 , 22:13:21                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 13:19:40                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func binsearch1(nums []int, target, low, high int) int {
	if len(nums) == 0 || nums[low] > target || nums[high] < target {
		return -1
	}
	mid := low + (high-low)>>1

	if target > nums[mid] {
		return binsearch1(nums, target, mid+1, high)
	} else if target < nums[mid] {
		return binsearch1(nums, target, low, mid-1)
	} else {
		return mid
	}
}
func search1(nums []int, target int) int {
	return binsearch1(nums, target, 0, len(nums)-1)
}

func search2(nums []int, target int) int {
	l := len(nums)
	low, high := 0, l-1
	if l == 0 {
		return -1
	}

	for low <= high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func search3(nums []int, target int) int {
	l := len(nums)
	low, high := 0, l-1
	if l == 0 {
		return -1
	}

	for low < high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if target == nums[low] {
		return low
	} else {
		return -1
	}
}

func search4(nums []int16, target int16) int16 {
	l := len(nums)
	low, high := 0, l-1
	if l == 0 {
		return -1
	}

	for low < high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if target == nums[low] {
		return int16(low)
	} else {
		return -1
	}
}

func main() {
	funcs := []func(numms []int, target int) int{
		// search1,
		// search2,
		// search3,
		// search4,
	}

	tests := []struct {
		nums   []int
		target int
		res    int
	}{
		// {nil, 55, -1},
		// {[]int{}, 5, -1},
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

	for i, f := range funcs {
		for _, test := range tests {
			res := f(test.nums, test.target)
			fmt.Printf("search%d(%v,%d) = %d: ", i+1, test.nums, test.target, res)

			if res != test.res {
				fmt.Println("failed")
			} else {
				fmt.Println("passed")
			}
		}
	}
}
