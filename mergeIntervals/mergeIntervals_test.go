/*
 * File: \mergeIntervals\mergeIntervals_test.go                                *
 * Project: leetcode                                                           *
 * Created At: Wednesday, 2022/03/30 , 00:44:43                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/03/30 , 01:41:46                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	funcs := map[string]func([][]int) [][]int{
		"merge1": merge1,
		"merge2": merge2,
		"merge3": merge3,
	}
	tests := []struct {
		intervals [][]int
		res       [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}},
		},
		{
			[][]int{{1, 4}, {0, 2}, {3, 5}}, [][]int{{0, 5}},
		},
	}

	for n, f := range funcs {
		for _, test := range tests {
			inter := make([][]int, len(test.intervals))
			copy(inter, test.intervals)

			if res := f(test.intervals); !reflect.DeepEqual(res, test.res) {
				t.Errorf("%s(%v) = %v,want %v\n", n, inter, res, test.res)
			}
		}
	}

}
