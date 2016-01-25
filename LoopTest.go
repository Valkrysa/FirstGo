package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 50.0 //50 selected at random, most any number would do
	for i := 0; i <= 10; i++ { //10 iterations as ordered by the tour guide
		z = z - ( (math.Pow(z, 2.0)-x)/(2*z) )
	}
	return z
}

//Learning material used from
//https://tour.golang.org/flowcontrol/8
func main() {
	fmt.Println(Sqrt(25))
	fmt.Println(math.Sqrt(25))
}