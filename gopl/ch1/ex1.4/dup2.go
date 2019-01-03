package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	countFiles := make(map[string][]string)
	if len(files) == 0 {
		countLines(os.Stdin, counts, countFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v \n", err)
				continue
			}
			countLines(f, counts, countFiles)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s%v\n", n, line, countFiles[line])
		}
	}

}

func countLines(f *os.File, counts map[string]int, countFiles map[string][]string) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !hasName(countFiles[text], name) {
			countFiles[text] = append(countFiles[text], name)
		}
	}
}

func hasName(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}

	return false
}
