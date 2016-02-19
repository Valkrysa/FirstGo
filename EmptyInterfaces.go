package main

import "fmt"

type Blank interface {
	M()
}

//Learning material used from
//https://tour.golang.org/methods/14
//https://tour.golang.org/methods/13
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	var b Blank
	describe(b)
	//b.M() //calling a method on a nil interface causes an error btw
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
