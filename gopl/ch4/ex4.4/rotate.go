package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s = rotate(s, 4)
	fmt.Println(s)
}

//rotate slice
func rotate(n []int, step int) []int {
	a := n[:step]
	b := n[step:]
	return append(b, a...)

}
