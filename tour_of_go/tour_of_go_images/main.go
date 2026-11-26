package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

// written by AI to see the result
func main() {
	m := Image{w: 256, h: 256}

	file, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, m); err != nil {
		panic(err)
	}
}
