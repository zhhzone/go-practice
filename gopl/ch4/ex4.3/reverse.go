package main

import (
	"fmt"
)

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	reverse1(&a)
	fmt.Printf("%v", a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]

	}
}

func reverse1(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]

	}

}
