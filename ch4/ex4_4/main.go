package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", rotate(s, 3))
}

// only supports positive value shifts
func rotate(s []int, shift int) []int {
	length := len(s)
	result := make([]int, length)
	for ix, val := range s {
		result[(ix + shift) % length] = val
	}
	return result
}
