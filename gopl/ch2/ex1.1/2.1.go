package main

import (
	"fmt"
	"github.com/zhhzone/go-practice/gopl/ch2/ex1.1/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.FToK(tempconv.CToF(tempconv.BoilingC)))
}
