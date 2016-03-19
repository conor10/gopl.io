// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) OpWith(t *IntSet, f func(uint64, uint64) uint64) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = f(s.words[i], tword)
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) UnionWith2(t *IntSet) {
	s.OpWith(t, s.or)
}

func (s *IntSet) IntersectWith(t *IntSet) {
	s.OpWith(t, s.and)
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	s.OpWith(t, s.xor)
}

func (s *IntSet) or(v1 uint64, v2 uint64) uint64 {
	return v1 | v2
}

func (s *IntSet) and(v1 uint64, v2 uint64) uint64 {
	return v1 & v2
}

func (s *IntSet) xor(v1 uint64, v2 uint64) uint64 {
	return v1 ^ v2
}


//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for i := word; i > 0; i >>= 1 {
			count += int(i & 1)
		}

	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	for idx, _ := range s.words {
		s.words[idx] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	dest := IntSet{make([]uint64, len(s.words))}
	copy(dest.words, s.words)
	return dest
}

func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}