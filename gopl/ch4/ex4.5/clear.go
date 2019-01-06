package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "one", "one", "three", "three"}
	fmt.Printf("%q\n", DeDuplicate(data))
}

func DeDuplicate(s []string) []string {
	var duplicate int
	length := len(s)
	for i := 0; i < length-1; i++ {
		j := i + 1
		fmt.Println(s[j])
		for ; j <= length-1-duplicate; j++ {
			fmt.Println(j)
			if s[i] != s[j] {
				break
			}
			duplicate++
		}

		copy(s[i:], s[j-1:])
	}
	fmt.Println(duplicate)
	return s[:len(s)-duplicate]
}
