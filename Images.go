package main

import (
	"fmt"
	"image"
)

//Learning material used from
//https://tour.golang.org/methods/24
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

/*
//Learning material used from
//https://tour.golang.org/methods/25
//requires running with golang.org
package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
	"math/rand"
)

func random(min, max int) uint8 {
	rand_value := rand.Intn(max - min) + min
	return uint8(rand_value)
}

type Image struct{
	w, h int
	color uint8
}

func (rect Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, rect.w, rect.h)
}

func (rect Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r Image) At(x, y int) color.Color {
	return color.RGBA{r.color+random(0, 255), r.color+uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(m)
}

 */