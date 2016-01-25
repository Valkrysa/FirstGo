package main

import "fmt"

func main() {
	fmt.Println("Beginning countdown.")
	defer fmt.Println("Blast off!")

	for i := 0; i <= 10; i++ {
		if i != 5 {
			defer fmt.Println(i)
		} else {
			defer fmt.Println("") // it is best not to say 5 when counting down
		}

	}
}
