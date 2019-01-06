package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "err style os.in", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	count := count(m, doc)

	for k, v := range count {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func count(m map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(m, c)
	}
	return m
}
