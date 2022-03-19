/*
 * File: \NumberOfBit1\hamingWeight.go                                         *
 * Project: leetcode                                                           *
 * Created At: Saturday, 2022/03/19 , 16:57:34                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/19 , 23:42:24                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func hamingWeight1(n uint) int {
	var pop uint

	for n != 0 {
		pop += n & 1
		n >>= 1
	}

	return int(pop)
}

func hamingWeight2(n uint) int {
	var pop int

	for n != 0 {
		pop++
		n &= (n - 1)
	}

	return pop
}

func hamingWeight3(n uint) int {
	a := n >> (0 * 8)
	fmt.Printf("n = %d,n>>0*8 = %d,byte(n>>0*8) = %d,pc[byte(n>>0*8)] = %d\n", n, a, byte(a), pc[byte(a)])
	b := n >> (1 * 8)
	fmt.Printf("n = %d,n>>1*8 = %d,byte(n>>1*8) = %d,pc[byte(n>>0*8)] = %d\n", n, b, byte(b), pc[byte(b)])
	c := n >> (2 * 8)
	fmt.Printf("n = %d,n>>2*8 = %d,byte(n>>2*8) = %d,pc[byte(n>>0*8)] = %d\n", n, c, byte(c), pc[byte(c)])
	d := n >> (3 * 8)
	fmt.Printf("n = %d,n>>3*8 = %d,byte(n>>3*8) = %d,pc[byte(n>>0*8)] = %d\n", n, d, byte(d), pc[byte(d)])
	e := n >> (4 * 8)
	f := n >> (5 * 8)
	g := n >> (6 * 8)
	h := n >> (7 * 8)

	return int(pc[byte(a)] + pc[byte(b)] + pc[byte(c)] + pc[byte(d)] + pc[byte(e)] + pc[byte(f)] + pc[byte(g)] + pc[byte(h)])
	// return int(pc[byte(n>>0*8)] + pc[byte(n>>1*8)] + pc[byte(n>>2*8)] + pc[byte(n>>3*8)] + pc[byte(n>>4*8)] + pc[byte(n>>5*8)] + pc[byte(n>>6*8)] + pc[byte(n>>7*8)])
}

func hamingWeight4(n uint) int {
	n = n&0x5555555555555555 + (n>>1)&0x5555555555555555
	n = n&0x3333333333333333 + (n>>2)&0x3333333333333333
	n = n&0x0F0F0F0F0F0F0F0F + (n>>4)&0x0F0F0F0F0F0F0F0F
	n = n&0x00FF00FF00FF00FF + (n>>8)&0x00FF00FF00FF00FF
	n = n&0x0000FFFF0000FFFF + (n>>16)&0x0000FFFF0000FFFF
	n = n&0x00000000FFFFFFFF + (n>>32)&0x00000000FFFFFFFF

	return int(n)
}
