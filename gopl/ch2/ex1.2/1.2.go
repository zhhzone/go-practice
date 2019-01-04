package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zhhzone/go-practice/gopl/ch2/ex1.2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

}
