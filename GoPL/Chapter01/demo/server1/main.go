// server1 是一个迷你回显服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HandleFunc 在默认服务多路复用器中注册给定模式的处理程序函数。
	// ServeMux 的文档解释了模式是如何匹配的。
	http.HandleFunc("/", handler) // 每个请求调用 hendler 函数
	// Fatal 等价于先调用 Print() 然后调用调用 os.exit(1)。
	// ListenAndServe 监听 TCP 网络地址 addr，然后调用带有处理程序
	// 的服务来处理传入连接的请求。接受的连接被配置为启用 TCP
	// keep-alives。handler 通常为 nil，在这种情况下，使用默认的服
	// 务多路复用器。ListenAndServe 总是返回一个非零错误。
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// HTTP 处理程序使用响应编写器(ResponseWriter)接口来构建 HTTP 响应。
// 在 Handler.ServeHTTP 方法返回后，不能使用响应编写器。
// Request 表示服务器接收到或客户端发送的 HTTP 请求。客户端和服务器
// 使用之间的字段语义略有不同。

// handler 回显请求 URL r 的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// 解释:
// 这个 Web 服务器程序之所以只有寥寥几行代码，是因为库函数做了大部
// 分工作。main 函数将一个处理函数和以"/"开头的 URL 连接接在一起，
// 代表所有的 URL 使用这个函数处理，然后启动服务器监听进入8000端口
// 处的请求。一个请求由一个 http.Request 类型的结构体表示，它包含
// 很多关联的域，其中一个是所请求的 URL。当一个请求到达时，它被转
// 交给处理函数，并从请求的 URL 中提取路径部分(/hello)，使用
// fmt.Fprintf 格式化，然后作为响应发送回去。
