package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, val := range os.Args {
		fmt.Print(idx)
		fmt.Println(": " + val)
	}
}
