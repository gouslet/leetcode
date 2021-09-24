package main

import "fmt"

func sortedSquares(nums []int) []int {
	l := len(nums) - 1
	res := make([]int, l+1, l+1)
	for i, j, k := 0, l, l; i <= j; {
		if a, b := nums[i]*nums[i], nums[j]*nums[j]; a <= b {
			res[k] = b
			j--
		} else {
			res[k] = a
			i++
		}
		k--
	}
	return res
}

func main() {
	nums := []int{-7, -3, 1, 3, 4}
	fmt.Printf("%v\n", sortedSquares(nums))
}
