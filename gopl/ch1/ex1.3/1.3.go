package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1(args []string) {
	var sep, s string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = ""
	}
}

func echo2(args []string) {
	var sep, s string
	for _, arg := range args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
}
func echo3(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)

}
