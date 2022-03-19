/*
 * File: \climbStairs\climbStairs.go                                           *
 * Project: leetcode                                                           *
 * Created At: Sunday, 2022/03/20 , 00:06:08                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/20 , 00:10:26                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

func climbStairs(n int) int {
	return 0
}

func square(n, m int) int {
	if m == 1 {
		return n
	}

	t := square(n, m>>1)
	if m&1 == 0 {
		return t * t
	} else {
		return t * t * m
	}
}
