package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.OpenFile("a.txt", os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Failed to OpenFile", err)
		return
	}

	n, err := fd.WriteString("hello world\n")
	if err != nil {
		fmt.Println("Failed to WriteString", err)
		return
	}
	fmt.Printf("write sucess %d bytes\n", n)

	fd.Seek(0, os.SEEK_SET)
	buf := make([]byte, 20)

	n, err = fd.Read(buf)
	os.Stdout.Write(buf)
	fd.Close()
}
