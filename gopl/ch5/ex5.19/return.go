package main

import (
	"fmt"
)

func main() {
	num, err := NoReturn()
	fmt.Println(num, err)
}

func NoReturn() (num int, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case bailout{}:
			num = 5
		default:
			panic(p)

		}
	}()
	panic(bailout{})
	return
}
