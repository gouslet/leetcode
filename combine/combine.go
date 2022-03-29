/*
 * File: \combine\combine.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Thursday, 2022/03/17 , 16:22:46                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/03/27 , 23:49:56                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

func combine(n, k int) [][]int {
	res := [][]int{}

	if n < 1 || k < 1 || n < k {
		return res
	}

	if n == k {
		var ints []int
		for i := 1; i <= n; i++ {
			ints = append(ints, i)
		}
		return [][]int{ints}
	}
	for i := n; i >= k; i-- {

		right := combine(i-1, k-1)

		if len(right) == 0 {
			res = append(res, []int{i})
		} else {
			for _, r := range right {

				res = append(res, append(r, i))
			}
		}

	}

	return res
}

func combine2(n, k int) [][]int {
	res := [][]int{}
	if n < 1 || k < 1 || n < k {
		return res
	}
	temp := []int{}

	var backtrack func(int)
	backtrack = func(cur int) {
		l := len(temp)
		if l+n-cur+1 < k {
			return
		}

		if l == k {
			res = append(res, append([]int{}, temp...))
			return
		}

		temp = append(temp, cur)
		backtrack(cur + 1)
		temp = temp[:len(temp)-1]
		backtrack(cur + 1)
	}

	backtrack(1)
	return res
}

// time O(n*2^n) 4ms 97.82% space O(n) 6336KB MB 93.01%
func combine3(n int, k int) [][]int {
	track := make([]int, 0, k)
	// var p func(int, int) int
	// p = func(a, b int) int {
	// 	res := 1
	// 	min := a - b
	// 	for a > min {
	// 		res *= a
	// 		a--
	// 	}

	// 	return res / 2
	// }

	// ans := make([][]int, 0, p(n, k))
	ans := make([][]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			ans = append(ans, append([]int{}, track...))
			return
		}
		for i := start; i <= n-(k-len(track))+1; i++ {
			// if len(track)+n-i+1 < k { //剪枝
			// 	return
			// }
			track = append(track, i)
			//选择
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return ans
}

// time O(n*2^n) 4ms 97.82% space O(k) 6340KB MB 92.04%
// combine4 字典序法实现组合型枚举
func combine4(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
