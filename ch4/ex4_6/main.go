package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

)

func main() {
	result := removeSpces([]byte("  this has     spaces    in  it .  "))
	fmt.Println(string(result))
}

func removeSpces(s []byte) []byte {
	for i := 0; i < len(s) - 1; {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			nextR, _ := utf8.DecodeRune(s[i+size:])
			if unicode.IsSpace(nextR) {
				copy(s[i:], s[i + size:])
				s = s[:len(s) - size]
			} else {
				i += size
			}
		} else {
			i += size
		}
	}
	return s
}
