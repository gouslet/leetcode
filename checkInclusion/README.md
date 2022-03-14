# 567. 字符串的排列

给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true；否则，返回 false。

换句话说，s1 的排列之一是 s2 的子串。

## 示例

### 示例 1：

输入：s1 = "ab" s2 = "eidbaooo"  
输出：true  
解释：s2 包含 s1 的排列之一 ("ba").

### 示例 2：

输入：s1= "ab" s2 = "eidboaoo"  
输出：false

## 提示：

- 1 <= s1.length, s2.length <= 104
- s1 和 s2 仅包含小写字母

## 解答

### 方法一

**解释**：
可用滑动窗口法：

- 因为 s1 和 s2 中仅含小写字母，所以可以用长度为 26 的数组 freq 来代替 map，
- 先统计 s1 中每个字母出现的频率，记录在 freq1 中
- 初始化 freq2 中每个元素的值为 0
- 增大窗口的右端 high，当`freq2[s2[high] - 'a'] + 1 > freq1[s2[high] - 'a']`时，将 low 向右移动至`s[high]`首次出现的位置，并将沿途的字母（不包括`s[high]`）出现的频率清零
- 当 high - low == s1 的长度时，返回 true
- 若 high 增长至 s2 右端时，若 high - low == s1 的长度，返回 true
- 否则返回 false

**代码**：

```go
func checkInclusion(s1 string, s2 string) bool {
	if s1 == "" || s2 == "nil" {
		return false
	}

	freq1 := [26]int{0}
	for _, c := range s1 {
		freq1[byte(c)-'a']++
	}
	freq2 := [26]int{0}

	low, high,i, l1, l2 := 0, 0, 0, len(s1), len(s2)
	for high <= l2 {
		if high-low == l1 {
			return true
		}
		if t := freq2[s2[high]-'a'] + 1; t > freq1[s2[high]-'a'] {
			for s2[i] != s2[high] {
				freq2[s2[i]-'a']--
				i++
			}
			i++
			low = i
		} else {
			freq2[s2[high]-'a']++
		}

		high++
	}

	// if high-low == l1 {
	// 	return true
	// }

	return false
}
```

### 方法二

**解释**：
可用频率比较法：

- 因为 s1 和 s2 中仅含小写字母，所以可以用长度为 26 的数组 freq 来统计 s1 和 s2 中的字母出现的频率，初始化 freq1 和 freq1 中每个元素的值为 0
- 当$i \in [0,len(s1))$时，统计 s1 中相应字符出现的频率，记录在 freq1 中，同时记录 s2 中相应位置的字符的频率在 freq2 中
- 当$i \in [len(s1),len(s2))$时，仅改变 freq2，
  - $freq2[s2[2]-'a']++$将 s2 中位置为 i 的字符的频率加 1
  - $freq2[s2[i-len(s1)]-' a']--$将 s2 中位置为 i-len(s1) 的字符的频率减 1
  - 若 freq1 == freq2，则返回 true
- 返回 false

**代码**：

```go
func checkInclusion(s1 string, s2 string) bool {
	if s1 == "" || s2 == "nil" {
		return false
	}

	freq1 := [26]int{0}
	for _, c := range s1 {
		freq1[byte(c)-'a']++
	}
	freq2 := [26]int{0}

	low, high,i, l1, l2 := 0, 0, 0, len(s1), len(s2)
	for high <= l2 {
		if high-low == l1 {
			return true
		}
		if t := freq2[s2[high]-'a'] + 1; t > freq1[s2[high]-'a'] {
			for s2[i] != s2[high] {
				freq2[s2[i]-'a']--
				i++
			}
			i++
			low = i
		} else {
			freq2[s2[high]-'a']++
		}

		high++
	}

	// if high-low == l1 {
	// 	return true
	// }

	return false
}
```
