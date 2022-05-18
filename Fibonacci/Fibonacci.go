/*
 * File: \Fibonacci\Fibonacci.go                                               *
 * Project: leetcode                                                           *
 * Created At: Tuesday, 2022/05/3 , 09:39:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/3 , 10:49:46                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "fmt"

func fib1(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		b %= (1e9 + 7)
		a, b = b, a+b
	}
	return a
}

func fib(n int) int {
	type matrix [2][2]int

	if n < 2 {
		return n
	}
	multiply := func(a, b matrix) matrix {
		var j int
		res := matrix{}
		for i := 0; i < 2; i++ {
			for j = 0; j < 2; j++ {
				res[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % (1e9 + 7)
			}
		}

		return res
	}

	var pow func(matrix, int) matrix

	pow = func(a matrix, n int) matrix {
		// fmt.Printf("pow(a,%d) = ", n)
		if n == 1 {
			// fmt.Println("a")
			return a
		}
		if n == 2 {
			// fmt.Println("a*a")
			return multiply(a, a)
		}
		if n%2 == 0 {
			// fmt.Printf("pow(pow(a,%d),%d)\n", n>>1, 2)
			return pow(pow(a, n/2), 2)
		}
		// fmt.Printf("multiply(pow(pow(a,%d),%d),a)\n", (n-1)>>1, 2)
		return multiply(pow(pow(a, (n-1)>>1), 2), a)
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)

	return res[0][0]
}

func main() {
	fmt.Println(fib(100))
}
