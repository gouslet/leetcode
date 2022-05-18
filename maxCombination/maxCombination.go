/*
 * File: \maxCombination\maxCombination.go                                     *
 * Project: leetcode                                                           *
 * Created At: Monday, 2022/05/16 , 03:09:19                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/16 , 03:32:04                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"sort"
	"strconv"
	"strings"
)

// func largestNumber(nums []int) string {
// 	strs := []string{}
// 	for _, n := range nums {
// 		strs = append(strs, strconv.Itoa(n))
// 	}
// 	sort.Slice(strs, func(i, j int) bool {
// 		a, _ := strconv.ParseInt(strs[i]+strs[j], 10, 64)
// 		b, _ := strconv.ParseInt(strs[j]+strs[i], 10, 64)

// 		return a > b
// 	})
// 	sb := strings.Builder{}
// 	for _, s := range strs {
// 		sb.WriteString(s)
// 	}
// 	res := sb.String()
// 	// res := strings.Join(strs, "")
// 	if b := strings.TrimLeft(res, "0"); b == "" {
// 		return "0"
// 	} else {
// 		return b
// 	}
// }

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		timesA, timesB := 10, 10
		for timesA <= a {
			timesA *= 10
		}

		for timesB <= b {
			timesB *= 10
		}

		return a*timesB+b > b*timesA+a
	})
	sb := strings.Builder{}
	for _, s := range nums {
		sb.WriteString(strconv.Itoa(s))
	}
	res := sb.String()
	// res := strings.Join(strs, "")
	if b := strings.TrimLeft(res, "0"); b == "" {
		return "0"
	} else {
		return b
	}
}
