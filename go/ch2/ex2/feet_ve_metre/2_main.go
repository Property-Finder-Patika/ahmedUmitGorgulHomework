package main

import "fmt"

type Feet float64

type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%gf", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func feetToM(f Feet) Meter {
	if f == 0.0 {
		return Meter(0.0)
	}
	return Meter(f / 3.28)
}

func MetreToF(m Meter) Feet {
	if m == 0.0 {
		return Feet(0.0)
	}
	return Feet(m * 3.28)
}
