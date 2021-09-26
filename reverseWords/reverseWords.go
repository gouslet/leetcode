package main

import (
	"fmt"
	"strings"
)

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

func main() {
	tests := []struct {
		s   string
		res string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, test := range tests {
		res := reverseWords(test.s)
		fmt.Printf("reverseWords(%q) = %q: ", test.s, res)
		if test.res != res {
			fmt.Println("failed")

		} else {
			fmt.Println("passed")
		}
	}
}
