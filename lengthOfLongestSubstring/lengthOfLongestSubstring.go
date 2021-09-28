package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
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

func main() {
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

	for _, test := range tests {
		maxLen := lengthOfLongestSubstring(test.s)
		fmt.Printf("lengthOfLongestSubstring(%q) = %d: ", test.s, maxLen)
		if maxLen != test.res {
			fmt.Println("failed")
		} else {
			fmt.Println("passed")
		}
	}
}
