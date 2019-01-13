package main

import (
	"io"
	"log"
	"net"
	"os"
)

type host struct {
	name    string
	address string
	port    string
}

func main() {
	//parse flag
	hosts := make([]host, 0)
	if len(os.Args) > 1 {
		for i, str := range os.Args {
			if i == 0 {
				continue
			}
			split := strings.Split(str, "=")
			name := split[0]
			split2 := strings.Split(split1[1], ":")
			address := split2[0]
			port := split2[1]
			hosts = append(hosts, host{name, address, port})
		}
	}

	// go func
}

func connect(port string) {
	conn, err := net.Dial("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer conn.Close()

	mustCopy(os.Stdout, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
