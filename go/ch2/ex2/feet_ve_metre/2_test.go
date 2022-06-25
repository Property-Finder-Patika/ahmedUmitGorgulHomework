package main

import "fmt"

func Example_length() {
	{
		f := MetreToF(10.0)
		fmt.Printf("%v\n", f.String())
		fmt.Printf("%g\n", feetToM(f))
		fmt.Printf("%g\n", feetToM(0))
		fmt.Printf("%g\n", MetreToF(0))
		fmt.Printf("%v\n", feetToM(3).String())
	}
}
