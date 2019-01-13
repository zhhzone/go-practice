package main

import (
	"fmt"
)

type Palindrome []byte

func (p Palindrome) Len() int {
	return len(p)
}

func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Palindrome) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func IsPalindrome(p Palindrome) bool {
	for i, j := 0, p.Len()-1; i < j; i, j = i+1, j-1 {
		if !p.Less(i, j) && !p.Less(j, i) {
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(Palindrome([]byte("absdkjljklsd"))))
	fmt.Println(IsPalindrome(Palindrome([]byte("1234554321"))))
}
