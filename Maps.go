package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var literal_map = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var literal_map_no_types = map[string]Vertex{
"Bell Labs": {40.68433, -74.39967},
"Google":    {37.42202, -122.08408},
}

//Learning material used from
//https://tour.golang.org/moretypes/16
//https://tour.golang.org/moretypes/17
//https://tour.golang.org/moretypes/18
//https://tour.golang.org/moretypes/19
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(literal_map)

	fmt.Println(literal_map_no_types)

	fmt.Println("===================")

	map_manipulation := make(map[string]int)

	map_manipulation["Answer"] = 42
	fmt.Println("The value:", map_manipulation["Answer"])

	map_manipulation["Answer"] = 48
	fmt.Println("The value:", map_manipulation["Answer"])

	delete(map_manipulation, "Answer")
	fmt.Println("The value:", map_manipulation["Answer"])

	v, ok := map_manipulation["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

/*
////Learning material used from
//https://tour.golang.org/moretypes/20
//GoTour needs to be installed locally
package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)

	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		word_count[ words[i] ] += 1
	}

	return word_count
}

func main() {
	wc.Test(WordCount)
}

 */