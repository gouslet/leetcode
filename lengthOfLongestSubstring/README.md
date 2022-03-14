# 3. 无重复字符的最长子串
给定一个字符串s，请你找出其中不含有重复字符的最长子串的长度。

## 示例
### 示例 1:

输入: s = "abcabcbb"  
输出: 3   
解释: 因为无重复字符的最长子串是"abc"，所以其长度为3。

### 示例 2:

输入: s = "bbbbb"  
输出: 1  
解释: 因为无重复字符的最长子串是"b"，所以其长度为 1。

### 示例 3:

输入: s = "pwwkew"  
输出: 3  
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

### 示例 4:

输入: s = ""  
输出: 0
 

## 提示：

- 0 <= s.length <= 5 * 104
- s由英文字母、数字、符号和空格组成

## 解答
### 方法一
**解释**：  
可采用滑动窗口法，
- 取窗口的左右端点分别为`low`和`high`，初始时`low`和`high`均为`0`，
- 用`map[string]int`类型的变量`exists`记录当前窗口中每个位置上的字符在整个字符串中的位置之后的位置（从`1`开始），
- 用`maxLen`记录最长无重复子串的长度值，初始为`0`
- 在`high`向右端滑行的过程中，
  - 如果`high`位置上的字符不在`map`中，则`high`继续向右移动，
  - 如果其存在于`map`中，则将`maxLen`值设置为当前的`high-low`值（窗口大小）与当前`maxLen`值中较大者，并将左端点移动至目前`s[high]`的值第一次出现的位置后面（在本程序中为`exists[s[high]]`）
  - 如果`high`到达字符串末端，则将`maxLen`值设置为当前的`high-low`值（窗口大小）与当前`maxLen`值中较大者

在整个过程中，`low-high`之间的字符无重复

**代码**：
```go
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	l := len(s)
	if l == 1 {
		return 1
	}
	maxLen := 0

	exists := make(map[byte]int)
	low, high := 0, 0
	for high < l {
		if i := exists[s[high]]; i != 0 {
			if diff := high - low; diff > maxLen {
				maxLen = diff
			}

			for low < i {
				exists[s[low]] = 0
				low++
			}

		} else if high == l-1 {
			if diff := high - low + 1; diff > maxLen {
				maxLen = diff
			}
		}
		exists[s[high]] = high + 1
		high++
	}
	return maxLen
}
```
具体见[lengthOfLongestSubstring](lengthOfLongestSubstring.go)