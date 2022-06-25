package tempconv

import "fmt"

type Celsius float64

type Fahrenheit float64

type Kelvin float64

const (
	AbsoluteZero_C Celsius = -273.15
	Freezing_C     Celsius = 0
	Boiling_C      Celsius = 100
	AbsoluteZero_K Kelvin  = 0
	Freezing_K     Kelvin  = 273.15
	Boiling_K      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZero_C) }

func KToC(k Kelvin) Celsius { return Celsius(k - Freezing_K) }
