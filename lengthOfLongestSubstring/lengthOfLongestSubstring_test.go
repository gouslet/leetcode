/*
 * File: \lengthOfLongestSubstring\lengthLongestSubstring.go                   *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/14 , 23:50:44                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/14 , 23:52:54                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "testing"

func TestLengthLongestSubstring(t *testing.T) {
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
			if maxLen := f(test.s); maxLen != test.res {
				t.Logf("%s(%q) = %d,want %d", fname, test.s, maxLen, test.res)
			}
		}
	}

}
