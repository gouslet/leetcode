/*
 * File: \rotate\main.go                                                       *
 * Project: leetcode                                                           *
 * Created At: Saturday, 2021/09/25 , 02:13:14                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/03/8 , 16:45:41                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

// rotate1 时间复杂度$O(n + k)$ 空间复杂度$O(n)$
func rotate1(nums []int, k int) {
	l := len(nums)
	if k <= 0 || l == 0 {
		return
	}
	nums_cpy := make([]int, l)
	for i := 0; i < l; i++ {
		nums_cpy[(i+k)%l] = nums[i]
	}
	copy(nums, nums_cpy)
}

// rotate2 时间复杂度$O(n + k)$ 空间复杂度$O(n)$
func rotate2(nums []int, k int) {
	l := len(nums)
	if k <= 0 || l == 0 {
		return
	}
	extra := make([]int, k)
	// fmt.Printf("%v\n", extra)
	k %= l
	fmt.Printf("%v\n", nums[l-k:l])
	copy(extra, nums[l-k:l])
	fmt.Printf("%v\n", nums)
	copy(nums[k:l], nums)
	fmt.Printf("%v\n", nums)
	copy(nums[0:k], extra)
	fmt.Printf("%v\n", nums)
}

// rotate3 时间复杂度$O(n)$ 空间复杂度$O(n)$
func rotate3(nums []int, k int) {
	l := len(nums)
	if k <= 0 || l == 0 {
		return
	}
	k %= l
	n := 0
	for i := l - 1; i >= l-k && i >= k; i-- {
		nums[i], nums[i-k] = nums[i-k], nums[i]
		n++
	}
	if n < k {
		rotate3(nums[:l-n], k-n)
	} else {
		rotate3(nums[:l-k], k)
	}
	// fmt.Printf("%v\n", nums)
}

// rotate4 时间复杂度$O(n)$ 空间复杂度$O(1)$
// 拆分成三步：
// 1.整体翻转数组 [7，6，5，4，3，2，1]
// 2.以k索引为界限 翻转前半段 [5，6，7，4，3，2，1]
// 3.再翻转后半段 [5，6，7，1，2，3，4]
func rotate4(nums []int, k int) {
	l := len(nums)
	if k <= 0 || l == 0 {
		return
	}
	k %= l
	// fmt.Printf("%v\n", nums)
	for i, j := 0, l-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	for i, j := 0, k-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	for i, j := k, l-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// rotate5 时间复杂度$O(n)$ 空间复杂度$O(1)$
// 将每一次替换形成一个闭环
func rotate5(nums []int, k int) {
	l := len(nums)
	if k <= 0 || l == 0 {
		return
	}
	k %= l
	// fmt.Printf("%v\n", nums)
	for start, count := 0, 0; count < l; start++ {
		// fmt.Printf("count = %d\n", count)
		cur := start

		pre := nums[cur]
		next := (cur + k) % l
		temp := nums[next]
		nums[next] = pre
		pre = temp
		cur = next
		count++
		for start != cur {
			next = (cur + k) % l
			temp = nums[next]
			nums[next] = pre
			pre = temp
			cur = next
			count++
		}
	}
}

func main() {
	rotates := []func(nums []int, k int){
		rotate1,
		rotate2,
		rotate3,
		rotate4,
		rotate5,
	}
	arrays := [][]int{
		{},
		{1, 2},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6},
	}
	distances := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	}

	for i, l := 0, len(rotates); i < l; i++ {
		for j, m := 0, len(arrays); j < m; j++ {
			for k, n := 0, len(distances); k < n; k++ {
				nums := make([]int, len(arrays[j]))
				copy(nums, arrays[j])
				fmt.Printf("rotate%d(%v,%d) = ", i+1, nums, k)
				rotates[i](nums, k)
				fmt.Printf("%v\n", nums)
			}
		}
	}
}
