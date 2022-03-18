/*
 * File: \combine\combine.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Thursday, 2022/03/17 , 16:22:46                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/03/18 , 17:06:49                                *
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

func combine2(n int, k int) (ans [][]int) {
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
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
