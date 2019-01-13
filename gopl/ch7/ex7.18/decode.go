package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	childrenStr := ""
	for _, child := range e.Children {
		childrenStr += child.String()
	}
	return "<" + e.Type.Local + ">" + childrenStr + "</" + e.Type.Local + ">"
}

func main() {
	input := "<A><B><C>hello</C><D>abc</D></B><C>world</C></A>"
	fmt.Println("Parsing input:", input)
	dec := xml.NewDecoder(strings.NewReader(input))
	var root *Element
	var stack []*Element
	var parent *Element
	depth := 0
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "7.18:%v\n", err)
			os.Exit(1)
		}
		if tok, ok := tok.(xml.StartElement); ok {
			newElement := Element{tok.Name, tok.Attr, []Node{}}
			if root == nil {
				root = &newElement
			} else {
				(*parent).Children = append((*parent).Children, &newElement)
			}
			fmt.Printf("%p\n", parent)
			stack = append(stack, &newElement)
			depth++
			parent = &newElement
		}
		if _, ok := tok.(xml.EndElement); ok {
			parent = stack[depth-1]
			depth--
			stack = stack[0:depth]
		}
		if tok, ok := tok.(xml.CharData); ok {
			(*parent).Children = append((*parent).Children, CharData(string(tok)))
		}
	}

	fmt.Println("\nOutPut:", root.String())
}
