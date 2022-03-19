# 77. 组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

# 示例
## 示例 1：

### 输入：
n = 4, k = 2
### 输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
## 示例 2：

### 输入：
n = 1, k = 1
### 输出：
[[1]]
 
## 提示：
- 1 <= n <= 20
- 1 <= k <= n

# 解答
## 解答一
### 解释
用递归方法
- 若`k < 1 || n < 1 || n < k`，则仅有一种组合，即为`[1,n]`
- 若`k == n`，则仅有一种组合，即为`[1,n]`
- 若`k<n`，则`k`个数的组合的第一个元素`i`可从`[k,n]`中选择，其余`k-1`个元素从`[1,i-1]`中选择
### 代码
```go
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
```
具体见[combine.go](combine.go)

## 解答二
### 解释
迭代方法
