// 为服务器添加功能很容易。有一个有用的扩展是一个特定的 URL，它返回
// 某种排序的状态。例如，这个版本的程序完成和回显服务器一样的事情，
// 但同时返回请求的数量；URL /count 请求返回到现在为止的计数，去掉
// /count 请求本身。

// server2 是一个迷你的回显和计数器服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler 回显请求的 URL 的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	// Lock 方法锁住 m。如果锁已在使用中，则调用 goroutine 会阻塞，
	// 直到互斥锁可用。
	mu.Lock()
	count++
	// Unlock 方法解开锁 m。如果 m 在进入 Unlock 时没有被锁定，这
	// 是一个运行时错误。锁定的互斥锁与特定的 goroutine 无关。允许
	// 一个 goroutine 锁定互斥锁，然后安排另一个 goroutine 解锁它。
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter 回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// 解释:
// 这个服务器有两个处理函数，通过请求的 URL 来决定哪一个被调用:
// 请求 /count 调用 counter，其他的调用 handler。以 / 结尾的处理
// 模式匹配所有含有这个前缀的 URL。在后台，对于每个传入的请求，服
// 务器在不同的 goroutine 中运行该处理函数，这样它可以同时处理多
// 个请求。然而，如果两个并发的请求试图同时更新计数值 count，它可
// 能会不一致地增加，程序会产生一个严重的竞态 bug。为了避免该问题，
// 必须确保最多只有一个 goroutine 在同一时间访问变量，这正是
// mu.Lock() 和 mu.Unlock() 调用的作用。
