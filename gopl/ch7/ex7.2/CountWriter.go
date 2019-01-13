package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCount struct {
	w     io.Writer
	count int64
}

func (c *ByteCount) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)
	return n, err
}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCount{w, 0}

	return &c, &c.count
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello world\n")
	fmt.Println(*count)
}
