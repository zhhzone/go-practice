package main

import (
	"testing"
)

/*
*go test -bench=.
 */

var (
	args = []string{"exec arg0 arg1 arg2 arg3"}
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}
