package main

import (
	"bufio"
	"fmt"
	"bytes"
)

type WordCounter int
func (c* WordCounter) Write(p []byte) (int, error) {
	written, err := Write(p, bufio.ScanWords)
	*c += WordCounter(written)
	return written, err
}

type LineCounter int
func (c* LineCounter) Write(p []byte) (int, error) {
	written, err := Write(p, bufio.ScanLines)
	*c += LineCounter(written)
	return written, err
}

func Write(p []byte, split bufio.SplitFunc) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(split)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func main() {
	var wc WordCounter
	var lc LineCounter

	text := "Some text with a number\nof words\nover a number\nof lines\n"
	wc.Write([]byte(text))
	lc.Write([]byte(text))

	fmt.Println(wc)
	fmt.Println(lc)
}