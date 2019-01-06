package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("学 而  不   思  则     罔")
	c := removeDupSpace(b)
	fmt.Println(string(c))
}

func removeDupSpace(b []byte) []byte {
	length := len(b)
	skipped := 0
	for i, v := range b {
		j := i + 1
		if unicode.IsSpace(rune(v)) {
			for ; j <= length-1-skipped; j++ {
				if !unicode.IsSpace(rune(b[j])) {
					break
				}
				skipped++
			}
			copy(b[i:], b[j-1:])
		}
	}
	fmt.Println(b[:length-skipped])
	return b[:length-skipped]
}
