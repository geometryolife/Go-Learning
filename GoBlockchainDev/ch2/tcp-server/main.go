package main

import (
	"fmt"
	"log"
	"net"
)

func handle_conn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connect ", conn.RemoteAddr())
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to Read", err)
			break
		}
		// fmt.Println(n, string(buf))
		fmt.Printf("size: %d, msg: %s", n, string(buf))
		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Failed to Write", err)
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to Listen", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept ", err)
		}
		go handle_conn(conn)
	}
}
