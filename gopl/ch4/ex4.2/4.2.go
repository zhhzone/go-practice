package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var m = flag.String("m", "256", "crypto num")

const (
	useSha256 = iota
	useSha384
	useSha512
)

func main() {
	flag.Parse()
	fmt.Println(*m)

	method := useSha256
	if len(os.Args) > 0 {
		switch *m {
		case "384":
			method = useSha384
		case "512":
			method = useSha512
		default:
			method = useSha256
		}
	}
	var input string
	fmt.Println("Enter text:")
	for {
		fmt.Scan(&input)
		if input == "" {
			continue
		}
		value := []byte(input)

		switch method {
		case useSha256:
			hash := sha256.Sum256(value)
			fmt.Println("%x", hash)
		case useSha384:
			hash := sha512.Sum384(value)
			fmt.Println("%x", hash)
		case useSha512:
			hash := sha512.Sum512(value)
			fmt.Println("%x", hash)
		}

	}
}

func PrintHash(hash []byte) {
	for _, v := range hash {
		fmt.Print("%X", v)
	}
	fmt.Println()
}
