package main

import (
	"fmt"
	"reflect"
)

func moveZeroes1(nums []int) {
	l := len(nums)
	if nums == nil || l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < l; j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

func moveZeroes2(nums []int) {
	l := len(nums)
	if nums == nil || l == 0 {
		return
	}
	j := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < l; j++ {
		nums[j] = 0
	}
}

func main() {
	funcs := []func(nums []int){
		moveZeroes1,
		moveZeroes2,
	}
	tests := []struct {
		nums []int
		res  []int
	}{
		// {
		// 	[]int{1, 1, 0, 1, 0},
		// 	[]int{1, 1, 1, 0, 0},
		// },
		// {
		// 	[]int{},
		// 	[]int{},
		// },
		// {
		// 	[]int{0, 1, 0, 1, 0},
		// 	[]int{1, 1, 0, 0, 0},
		// },
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
	}
	for i, l := 0, len(funcs); i < l; i++ {
		for _, test := range tests {
			n := test.nums
			fmt.Printf("moveZeros%d(%v) = ", i, n)
			funcs[i](n)
			fmt.Printf("%v ", n)
			if !reflect.DeepEqual(test.res, n) {
				fmt.Println("false")
			}
			fmt.Println("true")
		}
	}
}
