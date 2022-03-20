/*
 * File: \reverseBits\reverseBits.go                                           *
 * Project: leetcode                                                           *
 * Created At: Sunday, 2022/03/20 , 21:15:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/20 , 23:07:08                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

func singleNumber(nums []int) int {
	var res int
	for i := range nums {
		res ^= nums[i]
	}

	return res
}
