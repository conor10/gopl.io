package main

import "fmt"

func main() {
	// not sure how to print this
	fmt.Println(magic())
}

func magic() {
	defer recover()
	panic(1)
}
