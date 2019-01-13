package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topSort1(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items map[string][]string, index string)

	visitAll = func(items map[string][]string, index string) {
		if index != "" {
			for _, item := range items[index] {
				if !seen[item] {
					seen[item] = true
					visitAll(items, item)
					order = append(order, item)
				}
			}
		} else {
			for key, _ := range items {
				if !seen[key] {
					seen[key] = true
					visitAll(items, key)
				}
			}

		}
	}

	visitAll(m, "")
	return order
}
func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
