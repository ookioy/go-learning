package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// import "golang.org/x/tour"

// This is mine
func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)
	for y := range matrix {
		matrix[y] = make([]uint8, dx)
		for x := range matrix[y] {
			// You can change there a formula to see another image
			matrix[y][x] = uint8((x - y) / 2)
		}
	}
	return matrix
}

// func main() {
//     	pic.Show(Pic)
// }

// This part is written by an AI to see the result
func main() {
	const dx, dy = 256, 256
	data := Pic(dx, dy)

	img := image.NewGray(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.SetGray(x, y, color.Gray{Y: data[y][x]})
		}
	}

	f, err := os.Create("pic.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
