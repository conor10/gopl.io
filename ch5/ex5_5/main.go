package main

import (
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func main() {
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			data := strings.Split(c.Data, " ")
			words += len(data)
		} else if c.Type == html.ElementNode && c.Data == "img" {
			images++
		}
	}

	return
}