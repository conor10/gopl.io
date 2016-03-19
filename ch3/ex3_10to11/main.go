package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// Ex 3.10 & 3.11
// TODO: Add support for scientific notation (e.g. 1.25e10)
func comma(value string) string {
	var result bytes.Buffer

	if len(value) > 0 && value[0] == '-' {
		result.WriteByte('-')
		value = value[1:]
	}

	fraction := ""
	if len(value) > 2 && value[len(value) - 3] == '.' {
		fraction = value[len(value) - 2:]
		value = value[:len(value) - 3]
	}

	for idx, val := range value {
		if idx > 0 && (len(value) - idx ) % 3 == 0 {
			result.WriteByte(',')
		}
		result.WriteRune(val)
	}

	if fraction != "" {
		result.WriteByte('.')
		result.WriteString(fraction)
	}

	return result.String()
}


