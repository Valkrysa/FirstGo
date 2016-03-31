package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if err != nil {
		return 0, err
	}

	for i := range b {
		if b[i] >= 'A' && b[i] <= 'z' {
			if b[i] > 'a' && b[i] < 'a'+13 || b[i] > 'A' && b[i] < 'A'+13 {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		}
	}

	return n, err
}

//Learning material used from
//https://tour.golang.org/methods/23
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
