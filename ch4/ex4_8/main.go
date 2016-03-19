package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letters := make(map[rune]int)    // counts of Unicode characters
	digits := make(map[rune]int)
	symbols := make(map[rune]int)
	other := make(map[rune]int)

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			digits[r]++
		} else if unicode.IsLetter(r) {
			letters[r]++
		} else if unicode.IsSymbol(r) {
			symbols[r]++
		} else {
			other[r]++
		}
		utflen[n]++
	}

	print("letters", letters)
	print("digits", digits)
	print("symbols", symbols)
	print("other", other)

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func print(name string, counts map[rune]int) {
	fmt.Printf("%s\nrune\tcount\n", name)
	for c, n := range counts {
		fmt.Printf("%s\t%q\t%d\n", c, n)
	}
}

//!-
