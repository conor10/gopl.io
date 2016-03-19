package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	var args = os.Args

	var start1 = time.Now()
	join(args)
	fmt.Println("join: ", time.Now().Sub(start1))

	var start2 = time.Now()
	loop(args)
	fmt.Println("loop: ", time.Now().Sub(start2))
}

func join(args[] string) string {
	return strings.Join(args, " ")
}

func loop(args[] string) string {
	var result string
	for _, val := range args {
		result += val + " "
	}
	return result
}
