package main

import "fmt"

//Learning material used from
//https://tour.golang.org/methods/15
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic // to remove panic, simply return "ok" as well and it'll assume you're going to check "ok"
	fmt.Println(f)
}
