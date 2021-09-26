package main

import (
	"fmt"
	"reflect"
)

func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) < 2 || target < numbers[0] {
		return nil
	}

	for i, l := 0, len(numbers); i < l; i++ {
		if j := binsearch(numbers, target-numbers[i], i+1, l-1); j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

func binsearch(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1

	if nums[mid] < target {
		return binsearch(nums, target, mid+1, high)
	} else if nums[mid] > target {
		return binsearch(nums, target, low, mid-1)
	} else {
		return mid
	}
}

func main() {
	tests := []struct {
		numbers []int
		target  int
		res     []int
	}{
		{
			[]int{2, 7, 11, 15}, 9, []int{1, 2},
		},
		{
			[]int{2, 3, 4}, 6, []int{1, 3},
		},
		{
			[]int{-1, 0}, -1, []int{1, 2},
		},
		{
			[]int{5, 25, 75}, 100, []int{2, 3},
		},
	}

	for i, l := 0, len(tests); i < l; i++ {
		indexes := twoSum(tests[i].numbers, tests[i].target)
		fmt.Printf("twoSum(%v,%d) = %v: ", tests[i].numbers, tests[i].target, indexes)
		if !reflect.DeepEqual(tests[i].res, indexes) {
			fmt.Println("failed")
		} else {
			fmt.Println("passed")
		}
	}
}
