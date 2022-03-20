/*
 * File: \reverseBits\reverseBits.go                                           *
 * Project: leetcode                                                           *
 * Created At: Sunday, 2022/03/20 , 21:15:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/20 , 23:25:20                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

// 错误，高位为0的情况未能正确处理
// func reverseBits(num uint32) uint32 {
// 	var res uint32 = 0
// 	for num != 0 {
// 		res <<= 1
// 		res += num & 1
// 		num >>= 1
// 	}

// 	return res
// }

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res <<= 1
		res += 1 & (num >> i)
	}

	return res
}
