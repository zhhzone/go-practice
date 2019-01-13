package main

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

func (s *IntSet) Len() int {
	length := int(0)
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 8; j++ {
			length += int(pc[byte(word>>(j*8))])
		}

	}

	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
}

func (s *IntSet) Add(x ...int) {
	for _, y := range x {
		s.Add(y)
	}
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}
func (s *IntSet) Copy() *IntSet {
	var n IntSet
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return &n
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Remove(144)
	fmt.Println(x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Printf(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}
