/*
 * File: \reverseWords\reverseWords.go                                         *
 * Project: leetcode                                                           *
 * Created At: Sunday, 2021/09/26 , 21:31:20                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/03/10 , 15:57:10                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"strings"
)

func reverseWords1(s string) string {
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

func reverseWords2(s string) string {
	bytes := []byte(s)
	for i, j, l := 0, 0, len(bytes); i < l; i++ {
		var reverseString = func(low, high int) {
			for low < high {
				bytes[low], bytes[high] = bytes[high], bytes[low]
				low++
				high--
			}
		}

		if bytes[i] == ' ' {
			reverseString(j, i-1)
			j = i + 1
		} else if i == l-1 {
			reverseString(j, i)
			j = i + 1
		}

	}

	return string(bytes)

}

func main() {
	tests := []struct {
		s   string
		res string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, test := range tests {
		res := reverseWords2(test.s)
		fmt.Printf("reverseWords(%q) = %q: ", test.s, res)
		if test.res != res {
			fmt.Println("failed")

		} else {
			fmt.Println("passed")
		}
	}
}
