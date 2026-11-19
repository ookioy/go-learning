package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 1.0

	for {
		delta := (z*z - x) / (2 * z)
		z -= delta
		// Commonly at first iteration delta maybe negative
		if math.Abs(delta) < 0.0000000001 {
			break
		}
	}

	return z

}

func main() {

	var number float64 = 1423

	math_sqrt_result := math.Sqrt(number)
	func_sqrt_result := Sqrt(number)
	diff := math_sqrt_result - func_sqrt_result

	fmt.Println("math.sqrt: ", math_sqrt_result)
	fmt.Println("func sqrt: ", func_sqrt_result)
	fmt.Printf("The difference with math.Sqrt: %v \n", diff)

}
