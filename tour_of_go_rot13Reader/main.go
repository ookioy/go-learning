package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13r.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 65 && b[i] <= 90:
			if b[i]+13 > 90 {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		case b[i] >= 97 && b[i] <= 122:
			if b[i]+13 > 122 {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		default:
			continue
		}
	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
