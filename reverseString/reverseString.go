package main

import (
	"fmt"
	"reflect"
)

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	tests := []struct {
		s   []byte
		res []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, test := range tests {
		fmt.Printf("reverseString(%q) = ", test.s)
		reverseString(test.s)
		fmt.Printf("%q: ", test.s)
		if !reflect.DeepEqual(test.s, test.res) {
			fmt.Println("failed")
		} else {
			fmt.Println("passed")
		}
	}
}
