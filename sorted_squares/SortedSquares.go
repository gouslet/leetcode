/*
 * File: \sorted_squares\main.go                                               *
 * Project: leetcode                                                           *
 * Created At: Saturday, 2021/09/25 , 02:13:14                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/03/8 , 11:20:10                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func sortedSquares1(nums []int) []int {
	l := len(nums)
	res := make([]int, l, l)
	for i, j, k := 0, l-1, l-1; i <= j; {
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

func sortedSquares2(nums []int) []int {
	l := len(nums)
	res := make([]int, l, l)
	for i, j, k := 0, l-1, l-1; k >= 0; k-- {
		if a, b := nums[i]*nums[i], nums[j]*nums[j]; a <= b {
			res[k] = b
			j--
		} else {
			res[k] = a
			i++
		}
	}
	return res
}

func main() {
	nums := []int{-7, -3, 1, 3, 4}
	fmt.Printf("%v\n", sortedSquares1(nums))
	fmt.Printf("%v\n", sortedSquares2(nums))
}
