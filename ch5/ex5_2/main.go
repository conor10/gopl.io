package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for key, value := range visit(make(map[string]int), doc) {
		fmt.Printf("%s:\t%d\n", key, value)
	}
}

func visit(freq map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		freq[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		freq = visit(freq, c)
	}

	return freq
}