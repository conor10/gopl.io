package main

import "fmt"

func main() {
	s := []string{"one", "one", "two", "three", "one", "four", "four", "four", "five" }

	fmt.Printf("%s\n", eliminateDupes(s))
}

func eliminateDupes(s []string) []string {
	for i := 0; i < len(s) - 1; {
		fmt.Printf("%v\n", s)
		if s[i] == s[i+1] {
			copy(s[i:], s[i + 1:])
			s = s[:len(s) - 1]
		} else {
			i++
		}
	}
	return s[:len(s)]
}
