package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	l, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := 0; i < l; i++ {
		val := b[i]
		if val >= 'A' && val <= 'Z' {
			b[i] = (val-'A'+13)%26 + 'A'
		} else if val >= 'a' && val <= 'z' {
			b[i] = (val-'a'+13)%26 + 'a'
		}
	}

	return l, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
