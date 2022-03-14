/*
 * File: \checkInclusion\checkInclusion.go                                     *
 * Project: leetcode                                                           *
 * Created At: Tuesday, 2021/09/28 , 14:30:15                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/14 , 13:53:19                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

// 错误的，无法处理s1中字母重复出现的情况
// func checkInclusion1(s1 string, s2 string) bool {
// 	if s1 == "" || s2 == "nil" {
// 		return false
// 	}
// 	exists := make(map[byte]int)

// 	for _, c := range s1 {
// 		exists[byte(c)] = -1
// 	}
// 	low, high, l1, l2 := 0, 0, len(s1), len(s2)
// 	for high < l2 {
// 		if high-low == l1 {
// 			return true
// 		}

// 		if exists[s2[high]] != -1 {
// 			t := exists[s2[high]]
// 			if t != 0 {
// 				i := low
// 				for s2[i] != s2[high] {
// 					exists[s2[i]] = -1
// 					i++
// 				}
// 				low = i
// 			} else {
// 				low = high + 1
// 				for _, c := range s1 {
// 					exists[byte(c)] = -1
// 				}
// 			}
// 		} else {
// 			exists[s2[high]] = high + 1
// 		}
// 		high++
// 	}

// 	if high-low == l1 {
// 		return true
// 	}

// 	return false
// }

// func checkInclusion(s1 string, s2 string) bool {
// 	if s1 == "" || s2 == "nil" {
// 		return false
// 	}

// 	freq1 := [26]int{0}
// 	for _, c := range s1 {
// 		freq1[byte(c)-'a']++
// 	}
// 	freq2 := [26]int{0}

// 	low, high, l1, l2 := 0, 0, len(s1), len(s2)
// 	for high < l2 {
// 		if high-low == l1 {
// 			return true
// 		}
// 		if t := freq2[s2[high]-'a'] + 1; t > freq1[s2[high]-'a'] {
// 			i := low
// 			for s2[i] != s2[high] {
// 				freq2[s2[i]-'a']--
// 				i++
// 			}
// 			i++
// 			low = i
// 		} else {
// 			freq2[s2[high]-'a']++
// 		}

// 		high++
// 	}

// 	if high-low == l1 {
// 		return true
// 	}

// 	return false
// }

func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {                                
        return false
    }

    var cnt1, cnt2 [26]int

    for i, c := range s1 {
        cnt1[byte(c) - 'a']++
        cnt2[s2[i] - 'a']++
    }
    if cnt1 == cnt2 {
        return true
    }

    for i := len(s1); i < len(s2); i++ {
        cnt2[s2[i] - 'a']++
        cnt2[s2[i - len(s1)] - 'a']--
        if cnt1 == cnt2 {
            return true
        }
    }

    return false
}

func main() {
	tests := []struct {
		s1, s2 string
		res    bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaooo", false},
		{"adc", "dcda", true},
		{"ab", "ab", true},
		{"hello", "ooolleoooleh", false},
		{"xy", "abcxydefacb", true},
	}

	for _, test := range tests {
		res := checkInclusion(test.s1, test.s2)
		fmt.Printf("checkInclusion(%q,%q) = %t: ", test.s1, test.s2, res)
		if test.res != res {
			fmt.Println("failed")
		} else {
			fmt.Println("passed")
		}
	}
}
