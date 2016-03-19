package main
import (
	"fmt"
	"sort"
)

func main() {
	a := sortRunes([]rune("aba"))
	b := sortRunes([]rune("something"))
	c := sortRunes([]rune("cderedc"))

	fmt.Println("a", isPalindrom(a))
	fmt.Println("b", isPalindrom(b))
	fmt.Println("c", isPalindrom(c))
}

func isPalindrom(s sort.Interface) bool {
	length := s.Len()
	for i, j := 0, length - 1; i < length / 2; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
