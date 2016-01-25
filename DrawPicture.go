package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
	"time"
)

func random(min, max int) uint8 {
	rand_value := rand.Intn(max - min) + min
	return uint8(rand_value)
}

func Pic(dx, dy int) [][]uint8 {
	pic_storage := make([][]uint8, dy)

	for i, _ := range pic_storage {
		pic_storage[i] = make([]uint8, dx)
		for a, _ := range pic_storage {
			pic_storage[i][a] = random(0, 255)
		}
	}

	return pic_storage
}

////Learning material used from
//https://tour.golang.org/moretypes/15
//GoTour needs to be installed locally
func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	pic.Show(Pic)
}