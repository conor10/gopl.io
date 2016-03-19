package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}


func main() {
	a := sha256.Sum256([]byte{'x'})
	b := sha256.Sum256([]byte{'X'})

	compare(&a, &b)
}

func compare(aPtr *[32]byte, bPtr *[32]byte) {
	count := 0
	for ix, val := range aPtr {
		diff := val ^ bPtr[ix]
		count += int(pc[diff])
	}
	fmt.Printf("Bit count that differs: %d\n", count)
	fmt.Printf("%b\n%b\n", *aPtr, *bPtr)
}


