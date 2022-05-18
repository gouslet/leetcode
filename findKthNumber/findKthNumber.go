/*
 * File: \findKthNumber\findKthNumber.go                                       *
 * Project: leetcode                                                           *
 * Created At: Wednesday, 2022/05/18 , 12:36:42                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/18 , 21:19:32                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"sort"
)

func findKthNumber1(m, n, k int) int {
	l := k
	count := make([]int, k)
	var i, j, x int
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			x = (i + 1) * (j + 1)
			if x <= k {
				count[x-1]++
			}
		}
	}

	for i = 0; i < l; i++ {
		k -= count[i]
		if k <= 0 {
			return i + 1
		}
	}

	return 0
}

// func findKthNumber(m, n, k int) int {
// 	var i, j int

// 	var f func(count int, alo, ahi, blo, bhi int, x int) int
// 	var g func(alo, ahi, blo, bhi int, x int) int
// 	var h func(alo, ahi, blo, bhi int, x int) int

// 	f = func(count int, alo, ahi, blo, bhi, x int) int {
// 		mid := alo + (ahi-alo)>>1
// 		s1 := (mid + 1) * (mid + 2) / 2
// 		if s1 == x {
// 			max := mid + 1
// 			for i = 1; i <= mid+1; i++ {
// 				if i*(mid+1-i) > max {
// 					max = i * (mid + 1 - i)
// 				}
// 			}

// 			return max
// 		} else if s1 < x {
// 			return f(count-s1, x-s1)
// 		}
// 	}

// 	s, l := m, n
// 	if m > n {
// 		s = n
// 		l = m
// 	}
// 	count1 := s * (s + 1) / 2
// 	count2 := count1
// 	count3 := m*n - count1 - count2

// 	if k <= count1 {
// 		return f(count1, 0, s-1, 0, s-1, k)
// 	} else if k <= count1+count3 {
// 		return h(s, l-2, 1, l-s-1, k-count1)
// 	} else {
// 		return g(0, s-1, 0, n-s, k-count1-count3)
// 	}
// }

func findKthNumber2(m, n, k int) int {
	g := func(x int) int {
		res := x / n * n
		for i, hi := x/n+1, m+1; i < hi; i++ {
			res += x / i
		}

		return res
	}

	low, high := 1, m*n
	mid := low + (high-low)>>1
	b := g(mid)
	for low <= high {
		if b >= k {
			if g(mid-1) < k {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = low + (high-low)>>1
		b = g(mid)
	}

	return -1
}

func findKthNumber(m, n, k int) int {
	// 返回符合条件的最小i
	return sort.Search(m*n, func(i int) bool {
		count := i / n * n

		for j := i/n + 1; j <= m; j++ {
			count += i / j
		}

		return count >= k
	})
}

func main() {

	fmt.Println(findKthNumber(2, 3, 6))
	// fmt.Println(findKthNumber(3, 3, 5))
	fmt.Println(findKthNumber(9895, 28405, 100787757))
}
