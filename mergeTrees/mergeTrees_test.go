/*
 * File: \mergeTrees\mergeTrees_test.go                                        *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/14 , 19:18:31                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/14 , 21:09:54                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	tests := []struct {
	}{}

	for _, test := range tests {
		fmt.Println(test)
	}
}

func TestNewTreeFrom(t *testing.T) {
	tests := []struct {
		nums []int
		str  string
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			"1->2->3->4->5->6->7",
		},
	}

	for _, test := range tests {
		tree := NewTreeFrom(test.nums)
		if s := fmt.Sprintln(tree); s != test.str {
			t.Errorf("NewTreeFrom(%v) = %v,want %v", test.nums, s, test.str)
		}
	}
}
