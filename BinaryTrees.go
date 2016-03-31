package main

import (
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
}
