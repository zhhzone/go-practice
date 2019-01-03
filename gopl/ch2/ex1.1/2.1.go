package main

import (
	tempconv "./tempconv"
	"fmt"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.FToK(tempconv.CToF(tempconv.BoilingC)))
}
