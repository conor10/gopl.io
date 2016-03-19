package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	names := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		findDupes(os.Stdin, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			findDupes(f, names)
			f.Close()
		}
	}
	for line, files := range names {
		if files != "" {
			fmt.Printf("%s\t%s\n", files, line)
		}
	}
}

func findDupes(f *os.File, counts map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] += f.Name() + " "
	}
	// NOTE: ignoring potential errors from input.Err()
}
