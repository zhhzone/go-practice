package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	for index, arg := range os.Args[0:] {
		s = "index: " + strconv.Itoa(index) + " args:" + arg
		fmt.Println(s)
	}
}
