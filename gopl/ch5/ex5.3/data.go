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

	count(nil, doc)

}

func count(m []string, n *html.Node) {
	if n.Type == html.TextNode && n.Data != "script" {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(m, c)
	}
}
