/*
 * File: \rotate\RotateTest.go                                                 *
 * Project: leetcode                                                           *
 * Created At: Tuesday, 2022/03/8 , 12:38:28                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/03/8 , 16:48:59                                *
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

func TestRotate(t *testing.T) {
	tests := []struct {
		rotates map[string]func(nums []int, k int)
		nums    [][]int
		k_res   map[int][][]int
	}{
		{
			map[string]func(nums []int, k int){
				"rotate1": rotate1,
				"rotate2": rotate2,
				"rotate3": rotate3,
				"rotate4": rotate4,
				"rotate5": rotate5,
			},
			[][]int{
				{},
				{1, 2},
				{1, 2, 3, 4},
				{1, 2, 3, 4, 5, 6},
			},
			map[int][][]int{
				0: {
					{},
					{1, 2},
					{1, 2, 3, 4},
					{1, 2, 3, 4, 5, 6},
				},
				1: {
					{},
					{2, 1},
					{4, 1, 2, 3},
					{6, 1, 2, 3, 4, 5},
				},
				2: {
					{},
					{1, 2},
					{3, 4, 1, 2},
					{5, 6, 1, 2, 3, 4},
				},
				3: {
					{},
					{2, 1},
					{2, 3, 4, 1},
					{4, 5, 6, 1, 2, 3},
				},
				4: {
					{},
					{1, 2},
					{1, 2, 3, 4},
					{3, 4, 5, 6, 1, 2},
				},
				5: {
					{},
					{2, 1},
					{4, 1, 2, 3},
					{2, 3, 4, 5, 6, 1},
				},
				6: {
					{},
					{1, 2},
					{3, 4, 1, 2},
					{1, 2, 3, 4, 5, 6},
				},
				7: {
					{},
					{2, 1},
					{2, 3, 4, 1},
					{6, 1, 2, 3, 4, 5},
				},
				8: {
					{},
					{1, 2},
					{1, 2, 3, 4},
					{5, 6, 1, 2, 3, 4},
				},
			},
		},
	}

	for _, test := range tests {
		for s, r := range test.rotates {
			for j, l := 0, len(test.nums); j < l; j++ {
				ns := test.nums[j]

				for k, res := range test.k_res {
					n := make([]int, len(ns))
					copy(n, ns)
					if r(n, k); !reflect.DeepEqual(n, res[j]) {
						t.Errorf("%s(%v,%d) = %v,want %v\n", s, ns, k, n, res[j])
					}
				}
			}
		}
	}

}
