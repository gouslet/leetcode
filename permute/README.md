# 46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

 
# 示例
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
- 若`nums`长度为1，则返回包含nums的数组
- 遍历nums，返回其除去当前元素后的全排列，将当前元素加入每一个排列，即得完整数组的全排列
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