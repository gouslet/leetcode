/*
 * File: \permute\permute.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Friday, 2022/03/18 , 17:21:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 00:56:52                              *
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

func TestLetterCasePermutation(t *testing.T) {
	tests := []struct {
		s   string
		res []string
	}{
		// {"", []string{}},
		// {"z", []string{"z", "Z"}},
		// {"1z", []string{"1z", "1Z"}},
		{"Ab1z", []string{"ab1Z", "aB1Z", "ab1z", "aB1z", "Ab1z", "AB1z", "Ab1Z", "AB1Z"}},
	}

	for _, test := range tests {
		res := LetterCasePermutation(test.s)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("permute(%q) = %v,want %v", test.s, res, test.res)
		}
	}
}

func TestLetterCasePermutation2(t *testing.T) {
	tests := []struct {
		s   string
		res []string
	}{
		// {"", []string{}},
		// {"z", []string{"z", "Z"}},
		// {"1z", []string{"1z", "1Z"}},
		{"Ab1z", []string{"ab1Z", "aB1Z", "ab1z", "aB1z", "Ab1z", "AB1z", "Ab1Z", "AB1Z"}},
	}

	for _, test := range tests {
		res := LetterCasePermutation2(test.s)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("permute(%q) = %v,want %v", test.s, res, test.res)
		}
	}
}

func TestLetterCasePermutation3(t *testing.T) {
	tests := []struct {
		s   string
		res []string
	}{
		// {"", []string{}},
		// {"z", []string{"z", "Z"}},
		// {"1z", []string{"1z", "1Z"}},
		{"Ab1z", []string{"ab1Z", "aB1Z", "ab1z", "aB1z", "Ab1z", "AB1z", "Ab1Z", "AB1Z"}},
	}

	for _, test := range tests {
		res := LetterCasePermutation3(test.s)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("LetterCasePermutation3(%q) = %v,want %v", test.s, res, test.res)
		}
	}
}

func TestLetterCasePermutation4(t *testing.T) {
	tests := []struct {
		s   string
		res []string
	}{
		// {"", []string{}},
		// {"z", []string{"z", "Z"}},
		// {"1z", []string{"1z", "1Z"}},
		{"Ab1z", []string{"ab1Z", "aB1Z", "ab1z", "aB1z", "Ab1z", "AB1z", "Ab1Z", "AB1Z"}},
	}

	for _, test := range tests {
		res := LetterCasePermutation4(test.s)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("LetterCasePermutation4(%q) = %v,want %v", test.s, res, test.res)
		}
	}
}

func TestLetterCasePermutation5(t *testing.T) {
	tests := []struct {
		s   string
		res []string
	}{
		// {"", []string{}},
		// {"z", []string{"z", "Z"}},
		// {"1z", []string{"1z", "1Z"}},
		{"Ab1z", []string{"ab1Z", "aB1Z", "ab1z", "aB1z", "Ab1z", "AB1z", "Ab1Z", "AB1Z"}},
	}

	for _, test := range tests {
		res := LetterCasePermutation5(test.s)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("LetterCasePermutation5(%q) = %v,want %v", test.s, res, test.res)
		}
	}
}
