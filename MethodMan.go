package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func NonMethodAbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Learning material used from
//https://tour.golang.org/methods/1
//https://tour.golang.org/methods/2
//https://tour.golang.org/methods/3
//https://tour.golang.org/methods/4
//https://tour.golang.org/methods/5
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(NonMethodAbs(v))
	f := MyFloat(-4)
	fmt.Println(f.Abs())

	fmt.Println("===================")

	v.Scale(10)
	fmt.Println(v.Abs())

	fmt.Println("===================")

	v = Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(NonMethodAbs(v))
}
