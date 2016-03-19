// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			attrs := []string{tok.Name.Local}
			for _, attr := range tok.Attr {
				attrs = append(attrs, attr.Name.Local + "=" + attr.Value)
			}
			stack = append(stack, append(attrs)) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData, xml.Attr:
			if containsAll(stack, os.Args[1:]) {
				var msg []string
				for _, val := range stack {
					msg = append(msg, strings.Join(val, " "))
				}
				fmt.Printf("%s: %s\n", strings.Join(msg, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(xs [][]string, y []string) bool {
	for len(y) <= len(xs) {
		if len(y) == 0 {
			return true
		}
		if containsAny(xs[0], y[0]) {
			y = y[1:]
		}
		xs = xs[1:]
	}
	return false
}

// reports if y exists in xs
func containsAny(xs []string, y string) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

//!-
