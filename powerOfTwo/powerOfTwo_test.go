/*
 * File: \powerOfTwo\powerOfTwo_test.go                                        *
 * Project: leetcode                                                           *
 * Created At: Saturday, 2022/03/19 , 14:21:31                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 16:15:46                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "testing"

const (
	UINT_MAX = ^uint(0)
	UINT_MIN = 0
	INT_MAX  = int(UINT_MAX >> 1)
	INT_MIN  = -INT_MAX - 1
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n   int
		res bool
	}{
		{-1024, false},
		{-512, false},
		{-2, false},
		{-1, false},
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{INT_MAX, false},
		{INT_MIN, false},
	}

	for _, test := range tests {
		if res := isPowerOfTwo(test.n); res != test.res {
			t.Errorf("isPowerOfTwo(%d) = %t,want %t\n", test.n, res, test.res)
		}
	}
}
