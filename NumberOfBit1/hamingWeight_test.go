/*
 * File: \NumberOfBit1\hamingWeight_test.go                                    *
 * Project: leetcode                                                           *
 * Created At: Saturday, 2022/03/19 , 16:59:42                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 23:24:47                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"math"
	"strconv"
	"testing"
)

func TestHamingWeight(t *testing.T) {
	tests := []struct {
		n   uint
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{1024, 1},
		{math.MaxUint, len(strconv.FormatUint(math.MaxUint, 2))},
		{math.MaxUint >> 1, len(strconv.FormatUint(math.MaxUint>>1, 2))},
		{math.MaxUint >> 64, 0},
		{math.MaxUint >> 65, 0},
	}

	for _, test := range tests {
		if res := hamingWeight1(test.n); res != test.res {
			t.Errorf("hamingWeight(%d) = %d,want %d\n", test.n, res, test.res)
		}
	}
}

func TestHamingWeight2(t *testing.T) {
	tests := []struct {
		n   uint
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{1024, 1},
		{math.MaxUint, len(strconv.FormatUint(math.MaxUint, 2))},
		{math.MaxUint >> 1, len(strconv.FormatUint(math.MaxUint>>1, 2))},
		{math.MaxUint >> 64, 0},
		{math.MaxUint >> 65, 0},
	}

	for _, test := range tests {
		if res := hamingWeight2(test.n); res != test.res {
			t.Errorf("hamingWeight2(%d) = %d,want %d\n", test.n, res, test.res)
		}
	}
}

func TestHamingWeight3(t *testing.T) {
	tests := []struct {
		n   uint
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{1024, 1},
		{math.MaxUint, len(strconv.FormatUint(math.MaxUint, 2))},
		{math.MaxUint >> 1, len(strconv.FormatUint(math.MaxUint>>1, 2))},
		{math.MaxUint >> 64, 0},
		{math.MaxUint >> 65, 0},
	}

	for _, test := range tests {
		if res := hamingWeight3(test.n); res != test.res {
			t.Errorf("hamingWeight3(%d) = %d,want %d\n", test.n, res, test.res)
		}
	}
}

func TestHamingWeight4(t *testing.T) {
	tests := []struct {
		n   uint
		res int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{1024, 1},
		{math.MaxUint, len(strconv.FormatUint(math.MaxUint, 2))},
		{math.MaxUint >> 1, len(strconv.FormatUint(math.MaxUint>>1, 2))},
		{math.MaxUint >> 64, 0},
		{math.MaxUint >> 65, 0},
	}

	for _, test := range tests {
		if res := hamingWeight4(test.n); res != test.res {
			t.Errorf("hamingWeight4(%d) = %d,want %d\n", test.n, res, test.res)
		}
	}
}
