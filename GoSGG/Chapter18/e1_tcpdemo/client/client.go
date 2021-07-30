// 编写一个 TCP 客户端程序，能连接到服务端的 8888 端口，客户端可以
// 发送单行数据，发送成功后退出。
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client's Dial() err =", err)
		return
	}
	// fmt.Println("conn success =", conn)
	// 功能1: 客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入[终端]
	// 从终端读取一行用户输入，并准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ReadString() err =", err)
	}
	// 将 line 发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.Write() err =", err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据，并退出了\n", n)
}
