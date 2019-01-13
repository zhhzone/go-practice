package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if er != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	if tcpconn, ok := conn.(*net.TCPConn); ok {
		tcpconn.CloseWrite()
	}
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, r io.Reader) {
	if _, err := io.Copy(dst, r); err != nil {
		log.Fatal(err)
	}
}
