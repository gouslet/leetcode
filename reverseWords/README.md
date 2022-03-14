# 557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

## 示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
 
## 提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

## 解答
### 方法
**解释**：将string按空格拆分为子字符串数组，然后将每个子字符串逆转，最后组合为一个字符串  
**代码**：
```go
func reverseWords(s string) string {
	ss := strings.Fields(s)

	for i, str := range ss {
		ss[i] = reverseString(str)
	}

	return strings.Join(ss, " ")
}

func reverseString(s string) string {
	bytes := []byte(s)

	for i, j := 0, len(bytes)-1; i < j; {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
```
具体见[reverseWords.go](reverseWords.go)