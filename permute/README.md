# 46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

`	# 示例
## 示例 1：

### 输入：
nums = [1,2,3]
### 输出：
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

## 示例 2：

### 输入：
nums = [0,1]
### 输出：
[[0,1],[1,0]]

## 示例 3：

### 输入：
nums = [1]
### 输出：
[[1]]
 

# 提示：
- 1 <= nums.length <= 6
- -10 <= nums[i] <= 10
- nums 中的所有整数 互不相同

# 解答：
## 解答一
### 解释
用递归方法
- 遍历nums，返回其除去当前元素后的全排列，将当前元素加入每一个排列，即得完整数组的全排列
- 若`nums`长度为1，则返回包含nums的数组
### 代码
```go
func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	l := len(nums)
	if l == 1 {
		res = append(res, nums)
	}

	for i, n := range nums {
		sub := make([]int, l)
		copy(sub, nums)
		sub[i], sub[0] = sub[0], sub[i]

		subres := permute(sub[1:])
		for j := range subres {
			subres[j] = append(subres[j], n)
		}
		res = append(res, subres...)
	}
	return res
}
```
### 复杂度分析
**时间复杂度**：  
$O(n \times n!)$，其中 $n$ 为序列的长度。

- 算法的复杂度首先受 $permutation$ 的调用次数制约，$permutation$ 的调用次数为 $\sum_{k = 1}^{n}{P(n, k)}$次，其中 $P(n, k) = \frac{n!}{(n - k)!} = n (n - 1) ··· (n - k + 1)$，该式被称作 $n$ 的 $k$ - 排列，或者部分排列。  
而 
	$
		\sum_{k = 1}^{n}{P(n, k)} 
		= n! + \frac{n!}{1!} + \frac{n!}{2!} + \frac{n!}{3!} + ··· + \frac{n!}{(n-1)!} 
		< 2n! + \frac{n!}{2} + \frac{n!}{2^2} + \frac{n!}{2^{n-2}} 
		< 3n!
	$
这说明 $permutation$ 的调用次数是 $O(n!)$ 的。

- 而对于 permutation 调用的每个叶结点（共 $n!$ 个），需要将当前答案使用 $O(n)$ 的时间复制到答案数组中，相乘得时间复杂度为 $O(n \times n!)$。

	因此时间复杂度为 $O(n \times n!)$。

**空间复杂度**：  
$O(n)$，其中 $n$ 为序列的长度。

除答案数组以外，递归函数在递归过程中需要为每一层递归函数分配栈空间，所以这里需要额外的空间且该空间取决于递归的深度，这里可知递归调用深度为 $O(n)$。

## 解答二
### 解释
这个问题可以看作有n个排列成一行的空格，需要从左往右依此填入题目给定的n个数，每个数只能使用一次。

回溯法,有递归函数permutation
- 获取nums长度为l,向permutation传一个当前需要确定的元素位置cur
- 如果$cur==l$，说明已经填完了$n$个位置，找到了一个可行的解，将$nums$复制 到答案数组中，递归结束。
- 如果$cur<l$，考虑这第$cur$个位置要填哪个数。
#### 标记数组法
- 根据题目要求肯定不能填已经填过的数，因此很容易想到的一个处理手段是我们定义一个标记数组$vis[]$来标记已经填过的数，
- 那么在填第$cur$个数的时候遍历$nums$，如果这个数没有被标记过，就尝试填入，并将其标记，继续尝试填下一个位置，即调用函数$permutation(cur + 1)$。
- 回溯的时候要撤销这一个位置填的数以及标记，并继续尝试其他没被标记过的数。
#### 交换元素法
- 可以将题目给定的数组$nums$划分成左右两个部分，左边的表示已经填过的数，右边表示待填的数，在回溯的时候只要动态维护这个数组即可。
- 具体来说，
  - 假设已经填到第$cur$个位置，那么$nums$数组中$[0,cur−1]$是已填过的数的集合，$[cur,l−1]$是待填的数的集合。
  - 尝试用$[cur,l−1]$里的数去填第$cur$个数，假设待填的数的下标为$i$，那么填完以后将第$i$个数和第$cur$个数交换，即能使得在填第$cur+1$个数的时候$nums$数组的$[0,cur]$部分为已填过的数，$[first+1,n−1]$为待填的数，
  - 回溯的时候交换回来即能完成撤销操作。

### 代码
```go
func permute2(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	l := len(nums)

	var permutation func(cur int)
	permutation = func(cur int) {
		if cur == l {
			output := make([]int, l)
			copy(output, nums)
			res = append(res, output)
			return
		}

		permutation(cur + 1)
		for i := cur + 1; i < l; i++ {
			nums[i], nums[cur] = nums[cur], nums[i]
			permutation(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}

	permutation(0)

	return res
}

```
### 复杂度分析
**时间复杂度**：  
$O(n \times n!)$，其中 $n$ 为序列的长度。

- 算法的复杂度首先受 $permutation$ 的调用次数制约，$permutation$ 的调用次数为 $\sum_{k = 1}^{n}{P(n, k)}$次，其中 $P(n, k) = \frac{n!}{(n - k)!} = n (n - 1) ··· (n - k + 1)$，该式被称作 $n$ 的 $k$ - 排列，或者部分排列。  
而 
	$\sum_{k = 1}^{n}{P(n, k)} 
		= n! + \frac{n!}{1!} + \frac{n!}{2!} + \frac{n!}{3!} + ··· + \frac{n!}{(n-1)!} 
		< 2n! + \frac{n!}{2} + \frac{n!}{2^2} + \frac{n!}{2^{n-2}} 
		< 3n!$
这说明 $permutation$ 的调用次数是 $O(n!)$ 的。

- 而对于 permutation 调用的每个叶结点（共 $n!$ 个），需要将当前答案使用 $O(n)$ 的时间复制到答案数组中，相乘得时间复杂度为 $O(n \times n!)$。

	因此时间复杂度为 $O(n \times n!)$。

**空间复杂度**：  
$O(n)$，其中 $n$ 为序列的长度。

除答案数组以外，递归函数在递归过程中需要为每一层递归函数分配栈空间，所以这里需要额外的空间且该空间取决于递归的深度，这里可知递归调用深度为 $O(n)$。

### 优化
- 生成空的答案数组时提供cap参数
- 复制nums数组时使用append([]int(nil),nums...)
- 让i从cur+1开始递增，i=cur时不用交换元素

### 代码
```go
// permute3 回溯法 优化
func permute3(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	l := len(nums)

	var fact func(int) int
	fact = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}

	res := make([][]int, 0, fact(l))

	var permutation func(cur int)
	permutation = func(cur int) {
		if cur == l {
			res = append(res, append([]int(nil), nums...))
			return
		}

		permutation(cur + 1)
		for i := cur+1; i < l; i++ {
			nums[i], nums[cur] = nums[cur], nums[i]
			permutation(cur + 1)
			nums[i], nums[cur] = nums[cur], nums[i]
		}
	}

	permutation(0)

	return res
}
```