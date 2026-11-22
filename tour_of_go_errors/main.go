package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for {
		delta := (z*z - x) / (2 * z)
		z -= delta
		// Commonly at first iteration delta maybe negative
		if math.Abs(delta) < 0.0000000001 {
			break
		}
	}

	return z, nil
}

func Print(res float64, err error) {
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
func main() {
	Print(Sqrt(2))
	Print(Sqrt(-2))
}
