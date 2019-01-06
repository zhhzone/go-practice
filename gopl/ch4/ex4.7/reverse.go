package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := []byte("思而不学则殆")
	fmt.Println(string(reverse(a)))
}

func reverse(s []byte) []byte {
	if utf8.RuneCount(s) == 1 {
		return s
	}

	_, n := utf8.DecodeRune(s)

	return append(reverse(s[n:]), s[:n]...)
}
