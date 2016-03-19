package main

import (
	"fmt"
)

func main() {
	array := [...]int{'1', '2', '3', '4', '5'}
	reverse(&array, len(array))
	fmt.Printf("%q\n", array)

}

func reverse(s *[5]int, length int) {
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
