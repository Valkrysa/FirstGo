package main

import (
	"fmt"
	"io"
	"strings"
)

//Learning material used from
//https://tour.golang.org/methods/21
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/*
//Learning material used from
//https://tour.golang.org/methods/22

//Couldn't do this one by myself, didn't understand the structure.
//After looking up the answer I understand better.

package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

type ErrEOF []byte

func (e ErrEOF) Error() string {
    return fmt.Sprintf("End of file")
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
    if b == nil || len(b) == 0 {
        err := ErrEOF(b)
        return 0, err
    }

    for i := 0; i < len(b); i++ {
        b[i] = 'A'
    }

    return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
 */
