package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitRead struct {
	r io.Reader
	n int64
}

func (r *LimitRead) Read(p []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.n {
		p = p[0:r.n]
	}
	n, err = r.r.Read(p)
	fmt.Print("%v", p)
	return
}

func Limit(r io.Reader, n int64) *LimitRead {
	return &LimitRead{r, n}
}

/**
* this is a test
 */
func main() {
	r := Limit(strings.NewReader("<html><body><h1>hello</h1>h1></body>body></html>html>aaaaaaaaaaaa"), 40)
	buffer := make([]byte, 1024)
	n, _ := r.Read(buffer)
	fmt.Printf(string(buffer), n)
}
