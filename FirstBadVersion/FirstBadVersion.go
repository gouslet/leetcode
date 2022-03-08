/*
 * File: \FirstBadVersion\FirstBadVersion.go                                   *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/7 , 13:44:06                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 19:29:46                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func main() {
	tests := []struct {
		tt       []bool
		bad      int
		firstBad int
	}{
		{[]bool{false, true}, 1, 1},
		{[]bool{false, true, true}, 2, 1},
		{[]bool{false, false, true, true, true}, 4, 2},
	}

	for _, test := range tests {
		// bad := test.bad
		var isBadVersion = func(i int) bool {
			return test.tt[i]
		}
		var firstBadVersion = func(n int) int {
			i, j := 1, n
			for i <= j {
				m := (i + j) / 2
				if isBadVersion(m) {
					if !isBadVersion(m - 1) {
						return m
					} else {
						j = m - 1
					}
				} else {
					if isBadVersion(m - 1) {
						return m
					} else {
						i = m + 1
					}
				}
			}
			return i
		}
		if fb := firstBadVersion(len(test.tt)); fb != test.firstBad {
			fmt.Println("failed")
		}
	}

}
