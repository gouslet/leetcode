/*
 * File: \replaceString\replaceString.go                                       *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/21 , 18:50:55                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/21 , 19:46:43                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"bytes"
	"regexp"
	"strings"
)

//遍历每个字符，替换每个代表空格的rune
func replaceSpace1(s string) string {
	var sb strings.Builder

	for _, ch := range s {
		if ch == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(ch)
		}
	}

	return sb.String()
}

//用标准库strings包中的ReplaceAll函数
func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

//用标准库strings包中的Replace函数
func replaceSpace3(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

//用标准库bytes包中的ReplaceAll函数
func replaceSpace4(s string) string {
	return string(bytes.ReplaceAll([]byte(s), []byte(" "), []byte("%20")))
}

//用标准库bytes包中的Replace函数
func replaceSpace5(s string) string {
	return string(bytes.Replace([]byte(s), []byte(" "), []byte("%20"), -1))
}

//用标准库regexp包中的ReplaceAllLiteralString函数
var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(` `)
}

func replaceSpace6(s string) string {
	return re.ReplaceAllLiteralString(s, "%20")
}
