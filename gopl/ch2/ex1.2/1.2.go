package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zhhzone/go-practice/gopl/ch2/ex1.2/tempconv"
)

func main() {
	if len(os.Args) <= 1 {
		for true {
			var arg string
			_, err := fmt.Scanf("%s", &arg)
			if err != nil {
				fmt.Println("Error %g", err)
			}
			turn(getFloat64Arg(arg))

		}
	}

	for _, arg := range os.Args[1:] {
		turn(getFloat64Arg(arg))
	}

}

func getFloat64Arg(arg string) float64 {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %n", err)
		os.Exit(1)
	}
	return t
}

func turn(arg float64) {
	f := tempconv.Fahrenheit(arg)
	c := tempconv.Celsius(arg)
	fmt.Printf("%s = %s, %s = %s",
		f, tempconv.FToC(f), c, tempconv.CToF(c))

	feet := tempconv.Feet(arg)
	meter := tempconv.Meter(arg)
	fmt.Printf("%s = %s, %s = %s", feet, tempconv.FeetToMeter(feet), meter, tempconv.MeterToFeet(meter))

	pound := tempconv.Pound(arg)
	kilogram := tempconv.Kilogram(arg)
	fmt.Printf("%s = %s, %s = %s", pound, tempconv.PoundToKilogram(pound), kilogram, tempconv.KilogramToPound(kilogram))
}
