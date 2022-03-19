# 784. 字母大小写全排列
给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。

返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
 

# 示例:
## 示例 1：

### 输入：
s = "a1b2"
### 输出：
["a1b2", "a1B2", "A1b2", "A1B2"]

## 示例 2:

### 输入: 
s = "3z4"
### 输出: 
["3z4","3Z4"]
 

# 提示:
- 1 <= s.length <= 12
- s 由小写英文字母、大写英文字母和数字组成

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