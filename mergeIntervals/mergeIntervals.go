/*
 * File: \mergeIntervals\mergeIntervals.go                                     *
 * Project: leetcode                                                           *
 * Created At: Wednesday, 2022/03/30 , 00:09:08                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/03/30 , 01:34:41                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "sort"

func merge3(intervals [][]int) [][]int {
	if intervals == nil {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t := intervals[0]
	res := [][]int{}

	for i, l := 1, len(intervals); i < l; i++ {
		if inter := intervals[i]; inter[0] > t[1] {
			res = append(res, t)
			t = inter
		} else if inter[1] > t[1] {
			t[1] = inter[1]
		}
	}
	res = append(res, t)

	return res
}

func merge2(intervals [][]int) [][]int {
	if intervals == nil {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t := intervals[0]
	for i := 1; i < len(intervals); {
		if inter := intervals[i]; inter[0] > t[1] {
			t = inter
			i++
		} else if inter[1] <= t[1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			t[1] = inter[1]
			intervals = append(intervals[:i], intervals[i+1:]...)
		}
	}

	return intervals
}

func merge1(intervals [][]int) [][]int {
	if intervals == nil {
		return nil
	}

	sort.Sort(secs(intervals))
	t := intervals[0]
	for i := 1; i < len(intervals); {
		if inter := intervals[i]; inter[0] > t[1] {
			t = inter
			i++
		} else if inter[1] <= t[1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			t[1] = inter[1]
			intervals = append(intervals[:i], intervals[i+1:]...)
		}
	}

	return intervals
}

type secs [][]int

func (a secs) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a secs) Len() int {
	return len(a)
}

func (a secs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
