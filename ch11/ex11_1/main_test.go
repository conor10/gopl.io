package main

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	type output struct{
		counts map[rune]int
		utflen [utf8.UTFMax + 1]int
		invalid int
	}

	var tests = []struct {
		input string
		want output
	}{
		{
			"a bbb cc",
			output{
				map[rune]int{'a': 1, 'b': 3, 'c': 2, ' ': 2},
				[utf8.UTFMax + 1]int{0, 8, 0, 0, 0},
				0,
			},
		},
		{
			"a \u4e16\u754c cc",
			output{
				map[rune]int{'a': 1, '\u4e16': 1, '\u754c': 1, 'c': 2, ' ': 2},
				[utf8.UTFMax + 1]int{0, 5, 0, 2, 0},
				0,
			},
		},
	}

	for _, test := range tests {
		input = new(bytes.Buffer)
		input.(*bytes.Buffer).WriteString(test.input)
		counts, utflen, invalid := charCount()
		result := output{counts, utflen, invalid}
		if !mapEqual(result.counts, test.want.counts) ||
			result.invalid != test.want.invalid ||
			result.utflen != test.want.utflen {
			t.Errorf("charCount(%q) = %v", test.input, result)
		}
	}
}

func mapEqual(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
