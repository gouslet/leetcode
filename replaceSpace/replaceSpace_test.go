/*
 * File: \replaceString\replaceString_test.go                                  *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/03/21 , 18:54:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/21 , 19:41:36                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "testing"

func TestReplaceSpace(t *testing.T) {
	funcs := map[string]func(s string) string{
		"replaceSpace1": replaceSpace1,
		"replaceSpace2": replaceSpace2,
		"replaceSpace3": replaceSpace3,
		"replaceSpace4": replaceSpace4,
		"replaceSpace5": replaceSpace5,
		"replaceSpace6": replaceSpace6,
	}

	tests := []struct {
		s   string
		res string
	}{
		{"", ""},
		{" ", "%20"},
		{"	", "	"}, //\t
		{"  ", "%20%20"}, //space+space
		{"	 ", "	%20"}, //\t+space
		{" 	", "%20	"}, //space+\t
		{"% ", "%%20"},   //%+space
		{"%% ", "%%%20"}, //%+space
		{"We are happy", "We%20are%20happy"},
	}
	for n, f := range funcs {
		for _, test := range tests {
			if res := f(test.s); res != test.res {
				t.Errorf("%s(%q) = %q,want %q\n", n, test.s, res, test.res)
			}
		}
	}

}
