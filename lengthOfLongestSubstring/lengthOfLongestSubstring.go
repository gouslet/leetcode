/*
 * File: \lengthOfLongestSubstring\lengthOfLongestSubstring.go                 *
 * Project: leetcode                                                           *
 * Created At: Tuesday, 2021/09/28 , 00:49:48                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/14 , 23:50:06                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	if s == "" {
		return 0
	}

	maxLen := 0

	exists := make(map[byte]int)

	for low, high, l := 0, 0, len(s); high < l; high++ {
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
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	if s == "" {
		return 0
	}

	maxLen := 0

	exists := make(map[byte]int)

	for low, high, l := 0, 0, len(s); high < l; high++ {
		if i, ok := exists[s[high]]; ok {
			if diff := high - low; diff > maxLen {
				maxLen = diff
			}

			for low < i {
				delete(exists, s[high])
				low++
			}

		} else if high == l-1 {
			if diff := high - low + 1; diff > maxLen {
				maxLen = diff
			}
		}
		exists[s[high]] = high + 1
	}
	return maxLen
}

func main() {

	funcs := map[string]func(s string) int{
		"lengthOfLongestSubstring1": lengthOfLongestSubstring1,
		"lengthOfLongestSubstring2": lengthOfLongestSubstring2,
	}
	tests := []struct {
		s   string
		res int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"bbbbbc", 2},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
	}

	for fname, f := range funcs {
		for _, test := range tests {
			maxLen := f(test.s)
			fmt.Printf("%s(%q) = %d: ", fname, test.s, maxLen)
			if maxLen != test.res {
				fmt.Println("failed")
			} else {
				fmt.Println("passed")
			}
		}
	}

}
