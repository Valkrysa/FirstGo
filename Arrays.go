package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

//Learning material used from
//https://tour.golang.org/moretypes/6
//https://tour.golang.org/moretypes/7
//https://tour.golang.org/moretypes/9
//https://tour.golang.org/moretypes/10
//https://tour.golang.org/moretypes/11
func main() {
	var t [2]string
	t[0] = "Hello"
	t[1] = "World"
	fmt.Println(t[0], t[1])
	fmt.Println(t)

	fmt.Println("===================")

	//This is a slice
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	fmt.Println("===================")

	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])

	fmt.Println("===================")

	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	fmt.Println("===================")

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
