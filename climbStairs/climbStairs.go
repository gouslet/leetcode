/*
 * File: \climbStairs\climbStairs.go                                           *
 * Project: leetcode                                                           *
 * Created At: Sunday, 2022/03/20 , 00:06:08                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/20 , 23:31:03                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func climbStairs(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
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

func main() {
	fmt.Println(climbStairs(5))
}
