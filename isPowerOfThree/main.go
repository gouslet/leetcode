package main

import "fmt"

func isPowerOfThree1(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 || n%3 != 0 {
		return false
	}
	return isPowerOfThree1(n / 3)
}

// int型整数中最大的3的幂次方数为^uint(0)>>1=4052555153018976267
func isPowerOfThree2(n int) bool {
	if n > 0 && 4052555153018976267%n == 0 {
		return true
	}
	return false
}

func main() {
	funcs := []func(n int) bool{
		isPowerOfThree1,
		isPowerOfThree2,
	}
	tests := []struct {
		n   int
		res bool
	}{
		{
			-3, false,
		},
		{
			-2, false,
		},
		{
			-1, false,
		},
		{
			0, false,
		},
		{
			1, true,
		},
		{
			2, false,
		},
		{
			3, true,
		},
	}
	for i, l := 0, len(funcs); i < l; i++ {
		for _, test := range tests {
			b := funcs[i](test.n)
			fmt.Printf("isPowerOfThree%d(%d) = %t:", i+1, test.n, b)
			if b == test.res {
				fmt.Println("passed")
			} else {
				fmt.Println("failed")
			}
		}
	}

}
