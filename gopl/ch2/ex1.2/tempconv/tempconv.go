package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

const (
	FreezingC     Celsius = 0
	AbsoluteZeroC Celsius = -273.15
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g feet", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g meter", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g pound", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}
