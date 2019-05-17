package main

import (
	"image"
	"image/color"

	"github.com/rezilience/readytogo/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 50)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 0, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
