package main

import (
	"fmt"
	"math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Sqrt(x float64) float64 {
	z := 50.0 //50 selected at random, most any number would do
	for i := 0; i <= 10; i++ { //10 iterations as ordered by the tour guide
		z = z - ( (math.Pow(z, 2.0)-x)/(2*z) )
	}
	return z
}

//Learning material used from
//https://tour.golang.org/flowcontrol/8
//https://tour.golang.org/moretypes/13
func main() {
	fmt.Println(Sqrt(25))
	fmt.Println(math.Sqrt(25))

	fmt.Println("===================")

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("===================")

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}