package main

import (
	"fmt"
	"io"
)

type NewReader string

func (n *NewReader) Read(p []byte) (n int, err error) {
	n = len(*r)
	copy(p, []byte(*r))
	err = io.EOF
	return
}

func NewReader(str string) *NewReader {
	var s NewReader
	s = NewReader(str)
	return &s
}

func main()
