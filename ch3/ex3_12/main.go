package main

import (
	"fmt"
	"os"
)

func main() {
	if isAnagram(os.Args[1]) {
		fmt.Println("It is an anagram")
	} else {
		fmt.Println("It is not an anagram")
	}
}

func isAnagram(input string) bool {
	result := ""
	for _, val := range input {
		result = string(val) + result
	}
	return input == result
}
