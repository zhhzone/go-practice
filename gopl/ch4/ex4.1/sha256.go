package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println("diff %d bite", DiffByte(c1, c2))
}

func DiffByte(hash1, hash2 [32]byte) int {
	count := 0
	for i := uint(0); i < 32; i++ {
		byte1 := hash1[i]
		byte2 := hash2[i]

		diff := byte1 ^ byte2
		count += int(pc[diff])

	}
	return count
}
