package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2  = &Vertex{1, 2} // has type *Vertex
)

//Learning material used from
//https://tour.golang.org/moretypes/3
//https://tour.golang.org/moretypes/4
//https://tour.golang.org/moretypes/5
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	v = Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p2, v2, v3)
}