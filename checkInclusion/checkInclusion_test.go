/*
 * File: \checkInclusion\checkInclusion_test.go                                *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/14 , 11:44:40                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/14 , 12:28:52                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1, s2 string
		res    bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaooo", false},
		{"adc", "dcda", true},
		{"ab", "ab", true},
		{"hello", "ooolleoooleh", false},
	}

	for _, test := range tests {
		res := checkInclusion(test.s1, test.s2)
		if test.res != res {
			t.Errorf("checkInclusion(%q,%q) = %t: ",test.s1, test.s2, res)
		}
	}	
}