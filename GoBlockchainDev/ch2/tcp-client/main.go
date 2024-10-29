package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Failed to Dial")
		return
	}
	defer conn.Close()
	buf := make([]byte, 256)
	for {
		n, _ := os.Stdin.Read(buf)
		conn.Write(buf[:n])
		n, _ = conn.Read(buf)
		os.Stdout.Write(buf[:n])
	}
}
