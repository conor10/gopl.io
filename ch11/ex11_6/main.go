// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// ex2.3
func PopCountByLooping(x uint64) int {
	var sum int = 0
	for i := uint(0); i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// ex2.4
func PopCountBitwise(x uint64) int {
	var sum int = 0
	for i := x; i > 0; i >>= 1 {
		sum += int(i & 1)
	}
	return sum
}

// ex2.5
func PopCountRightmostBit(x uint64) int {
	var sum int = 0
	for i := x; i > 0; i &= (i-1) {
		sum++
	}
	return sum
}

//!-
