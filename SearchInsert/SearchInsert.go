/*
 * File: \SearchInsert\SearchInsert.go                                         *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/7 , 20:08:32                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 20:35:29                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

func SearchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] < target {
			if mi == hi || nums[mi+1] >= target {
				return mi + 1
			}
			lo = mi + 1
		} else if nums[mi] > target {
			if mi == 0 || nums[mi-1] < target {
				return mi
			}
			hi = mi - 1
		} else {
			return mi
		}
	}
	return hi
}
