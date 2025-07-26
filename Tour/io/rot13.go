package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i := 0; i < n; i++ {
		char := p[i]
		if char >= 'A' && char <= 'Z' {
			p[i] = ((char - 'A' + 13) % 26) + 'A'
		} else if char >= 'a' && char <= 'z' {
			p[i] = ((char - 'a' + 13) % 26) + 'a'
		}
	}
	return n, err
}

func rot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
