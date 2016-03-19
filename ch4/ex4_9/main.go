package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	freq := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		freq[scanner.Text()]++
	}

	for k, v := range freq {
		fmt.Printf("%s:\t%d\n", k, v)
	}
}
