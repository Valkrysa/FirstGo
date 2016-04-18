package main

import (
	"reflect"
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		go Walk(t.Left, ch)
		ch <- t.Value
		go Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	channel_one := make(chan int)
	channel_two := make(chan int)

	go Walk(t1, channel_one)
	go Walk(t2, channel_two)

	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for i := 0; i < 10; i++ {
		m1[<-channel_one] = true;
		m2[<-channel_two] = true;
	}

	if reflect.DeepEqual(m1, m2) {
		return true
	} else {
		return false
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

//Learning material used from
//https://tour.golang.org/concurrency/7
//https://tour.golang.org/concurrency/8
//This won't work if not run on the golang site.
func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
