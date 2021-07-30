// 编写一个 TCP 服务器程序，在 8888 端口监听，可以和多个客户端创建连接，
// 连接成功后，服务端接收数据并显示在终端上。

// 做网络 socket 开发时，net 包含有我们所需要的所有方法和函数。
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("服务器开始监听......")
	// tcp 表示所使用的网络协议，127.0.0.1:8888 表示在本地监听 8888 端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Listen() err =", err)
		return
	}
	defer listen.Close()

	// 循环等待客户端来连接
	for {
		// 等待客户端连接
		fmt.Println("等待客户端连接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err =", err)
		} else {
			fmt.Printf("Accept() suc conn = %v, 客户端的地址是: %v\n",
				conn, conn.RemoteAddr().String())
		}
		// 这里启动一个协程，为客户端服务
		go process(conn)
	}

	// fmt.Printf("listen suc =%v\n", listen)
}

func process(conn net.Conn) {
	// 这里循环接收客户端发送的数据
	defer conn.Close() // 关闭主函数中生成的 conn

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// conn.Read 等待客户端通过 conn.Write 发送信息，如果客户端没有
		// conn.Write 发送信息，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端 %s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从 conn 读取
		/*
			if err != nil {
				fmt.Println("Server Read err =", err)
				return
			}
		*/
		if err == io.EOF {
			fmt.Println("客户端退出!")
			return
		}
		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n])) // 客户端发送的信息已添加换行符
	}
}
