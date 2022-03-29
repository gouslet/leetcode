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
回溯法
在递归函数backtrack中确定cur位置的元素是否选取，然后求解cur+1位置

### 代码
```go
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
```

## 解答三
### 解释
假设一个方案从低到高的$k$个数分别是$\{a_0, a_1, \cdots, a_{k - 1}\}$，可以从低位向高位找到第一个$j$使得$a_{j} + 1 \neq a_{j + 1}$，出现在$a$序列中的数字在二进制数中对应的位置一定是$1$，即表示被选中，那么$a_{j} + 1 \neq a_{j + 1}$意味着$a_j$和$a_{j + 1}$对应的二进制位中间有$0$，即这两个$1$不连续。

把$a_j$对应的$1$向高位推送，也就对应着$a_j \leftarrow a_j + 1$，而对于$i \in [0, j - 1]$内所有的$a_i$把值恢复成$i + 1$，即对应这$j$个$1$被移动到了二进制数的最低$j$位。

用一个数组 $\rm temp$ 来存放$a$序列，一开始先把 $1$ 到 $k$ 按顺序存入这个数组，他们对应的下标是 $0$ 到 $k - 1$。为了计算的方便，需要在下标 $k$ 的位置放置一个哨兵 $n + 1$（思考题：为什么是 $n + 1$呢？）。

然后对这个 $\rm temp$ 序列按照这个规则进行变换，每次把前 $k$ 位（即除了最后一位哨兵）的元素形成的子数组加入答案。

每次变换的时候，把第一个 $a_{j} + 1 \neq a_{j + 1}$的 $j$ 找出，使 $a_j$自增 $1$，同时对 $i \in [0, j - 1]$的 $a_i$重新置数。如此循环，直到$\rm temp$中的所有元素为$n$内最大的 $k$个元素。

回过头看这个思考题，它是为了判断退出条件服务的。如何判断枚举到了终止条件呢？其实不是直接通过 $\rm temp$ 来判断的，会看每次找到的 $j$ 的位置，如果 $j = k$ 了，就说明 $[0, k - 1]$ 内的所有的数字是比第 $k$ 位小的最后 $k$ 个数字，这个时候找不到任何方案的字典序比当前方案大了，结束枚举。

### 代码
```go
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
```

