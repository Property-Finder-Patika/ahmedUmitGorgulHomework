package 2_1

import "fmt"

func Example_one() {
	{
		fmt.Printf("%g\n", Boiling_C-Freezing_C)
		boilingF := CToF(Boiling_C)
		fmt.Printf("%v\n", boiling_F.String())
		fmt.Printf("%g\n", boiling_F-CToF(Freezing_C))
		boilingK := CToK(Boiling_C)
		fmt.Printf("%g\n", boiling_K-CToK(Freezing_C)) 
	}
}

func Example_two() {
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

	c = KToC(Freezing_K)
	fmt.Println(Freezing_K.String())
	fmt.Println(c.String())       
	fmt.Println(c.String())       
}