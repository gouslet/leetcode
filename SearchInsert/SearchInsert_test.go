/*
 * File: \SearchInsert\SearchInsert_test.go                                    *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/7 , 20:27:09                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/7 , 20:32:08                                 *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums        []int
		target, res int
	}{
		{
			[]int{1, 3, 5, 6}, 7, 4,
		},
	}

	for _, test := range tests {
		if b := SearchInsert(test.nums, test.target); b != test.res {
			t.Errorf("SearchInsert(%v,%d) = %d,want %d", test.nums, test.target, b, test.res)
		}
	}
}
