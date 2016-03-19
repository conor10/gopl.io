package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("$foo foo $foo $foo foo foo foo $foo ", Reverse))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func expand(s string, f func(string) string) string {
	result := ""

	for _, word := range strings.Split(s, " ") {
		runeIdx := strings.IndexRune(word, '$')
		if runeIdx != -1 {
			result += f(word[runeIdx + 1:len(word)]) + " "
		} else {
			result += word + " "
		}
	}

	return result
}
