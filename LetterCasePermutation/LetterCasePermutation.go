/*
 * File: \permute\permute.go                                                   *
 * Project: leetcode                                                           *
 * Created At: Friday, 2022/03/18 , 17:21:33                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 17:25:56                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"strings"
)

func LetterCasePermutation(s string) []string {
	if s == "" {
		return []string{""}
	}
	res := []string{}
	n := s[0]
	ss := LetterCasePermutation(s[1:])

	var sb strings.Builder
	for _, sss := range ss {
		sb.Reset()
		sb.WriteByte(n)
		sb.WriteString(sss)

		res = append(res, sb.String())

		if 'a' <= n && n <= 'z' || 'A' <= n && n <= 'Z' {
			sb.Reset()
			sb.WriteByte(n ^ 32)
			sb.WriteString(sss)
			res = append(res, sb.String())
		}
	}

	return res
}

func LetterCasePermutation2(s string) []string {
	if s == "" {
		return []string{""}
	}
	r := []string{}

	var f func(data []byte) [][]byte

	f = func(data []byte) [][]byte {
		l := len(data)
		res := [][]byte{}
		if l == 0 {
			return res
		}
		n := data[0]
		if l == 1 {
			res = append(res, data)
			if 'a' <= n && n <= 'z' || 'A' <= n && n <= 'Z' {
				res = append(res, []byte{n ^ 32})
			}
		}
		for _, sss := range f(data[1:]) {
			bs := []byte{}
			bs = append(bs, data[0])
			bs = append(bs, sss...)

			res = append(res, bs)

			if 'a' <= n && n <= 'z' || 'A' <= n && n <= 'Z' {
				bs1 := []byte{n ^ 32}
				bs1 = append(bs1, sss...)
				res = append(res, bs1)
			}
		}
		return res
	}

	for _, bytes := range f([]byte(s)) {
		r = append(r, string(bytes))
	}
	return r
}

func LetterCasePermutation3(s string) []string {
	if s == "" {
		return []string{""}
	}
	r := [][]byte{{s[0]}}

	if c := s[0]; 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
		r = append(r, []byte{c ^ 32})
	}

	for i, l := 1, len(s); i < l; i++ {
		c := byte(s[i])
		for j, bytes := range r {
			r[j] = append(r[j], c)
			if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
				bs := make([]byte, i)
				copy(bs, bytes)
				r = append(r, append(bs, c^32))
			}
		}
	}
	res := []string{}
	for _, bytes := range r {
		res = append(res, string(bytes))
	}
	return res
}

func LetterCasePermutation5(s string) []string {
	var res []string
	ss := []byte(s)
	res = res[0:0]

	var permutation func(ss []byte, pos int)

	permutation = func(ss []byte, pos int) {
		if pos == len(ss) {
			tmp := string(ss)
			res = append(res, tmp)
			return
		}
		permutation(ss, pos+1)
		if ss[pos] >= 'A' {
			ss[pos] ^= 32
			permutation(ss, pos+1)
		}
	}

	permutation(ss, 0)
	return res
}

func LetterCasePermutation4(s string) []string {
	var res []string
	ss := []byte(s)
	l := len(s)
	queue := [][]byte{ss}
	for j := 0; j < len(queue); j++ {
		for i := range s {
			ss = queue[j]
			if s[i] >= 'A' {
				bs := make([]byte, l)
				copy(bs, queue[j])
				queue = append(queue, bs)
				ss[i] ^= 32
			}
		}
	}
	for _, bytes := range queue {
		res = append(res, string(bytes))
	}
	return res
}
