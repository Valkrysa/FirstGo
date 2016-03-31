package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//Learning material used from
//https://tour.golang.org/concurrency/1
func main() {
	go say("world")
	say("hello")

	//Learning material used from
	//https://tour.golang.org/concurrency/2
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	
	fmt.Println(x, y, x+y)

	//Learning material used from
	//https://tour.golang.org/concurrency/3
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//buffered channel, if you were to try to add one more item you'd error out.
}