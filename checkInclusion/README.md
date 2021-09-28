# 567. 字符串的排列
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

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
- 因为s1和s2中仅含小写字母，所以可以用长度为26的数组freq来代替map，
- 先统计s1中每个字母出现的频率，记录在freq1中
- 初始化freq2中每个元素的值为0
- 增大窗口的右端high，当`freq2[s2[high] - 'a'] + 1 > freq1[s2[high] - 'a']`时，将low向右移动至`s[high]`首次出现的位置，并将沿途的字母（不包括`s[high]`）出现的频率清零
- 当high - low == s1的长度时，返回true
- 若high增长至s2右端时，若high - low == s1的长度，返回 true
- 否则返回false
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

	low, high, l1, l2 := 0, 0, len(s1), len(s2)
	for high < l2 {
		if high-low == l1 {
			return true
		}
		if t := freq2[s2[high]-'a'] + 1; t > freq1[s2[high]-'a'] {
			i := low
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

	if high-low == l1 {
		return true
	}

	return false
}
```