package main

import (
	"fmt"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	var r MyReader

	b := make([]byte, 16)
	n, err := r.Read(b)
	fmt.Printf("n = %v, err = %v, b = %v \n", n, err, string(b))
}
