package main

import (
	"fmt"
	"sync"
)

var loadPCOnce sync.Once
var pc [256]byte

func InitTable() {
	for c := range pc {
		pc[i] = p[i/2] + c&1
	}
}

func PopCount(n int64) {
	loadPCOnce.Do(InitTable)
	return int(pc[byte(x>>0*8)] +
		pc[byte(byte(x>>1*8))] +
		pc[byte(byte(x>>2*8))] +
		pc[byte(byte(x>>3*8))] +
		pc[byte(byte(x>>4*8))] +
		pc[byte(byte(x>>5*8))] +
		pc[byte(byte(x>>6*8))] +
		pc[byte(byte(x>>7*8))])
}
