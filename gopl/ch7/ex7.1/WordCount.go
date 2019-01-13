package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WorldCounter int

func (w *WorldCounter) Count(p []byte) (int, error) {
	doc := bufio.NewScanner(strings.NewReader(string(p)))

	doc.Split(bufio.ScanWords)
	c := 0
	for doc.Scan() {
		c++
	}
	*w += WorldCounter(c)
	return c, nil
}

type LineCounter int

func (w *LineCounter) Line(p []byte) (int, error) {
	doc := bufio.NewScanner(strings.NewReader(string(p)))
	c := 0
	for doc.Scan() {
		c++
	}

	*w += LineCounter(c)
	return c, nil
}

func main() {
	var wc WorldCounter
	wc.Count([]byte("hello world! 大撒比"))
	fmt.Println(wc)
	var line LineCounter
	line.Line([]byte(`hello
sssl
sssl
sssl
`))
	fmt.Println(line)

}
