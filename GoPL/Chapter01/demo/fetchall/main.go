// 1. goroutine 是一个并发执行的函数。
// 2. 通道是一种允许某一例程向另一例程传递指定类型的值的通信机制。
// 3. 对于每个命令行参数，go 语句在第一轮循环中启动一个新的 goroutine，
// 它异步调用 fetch 来使用 http.Get 获取 URL 内容。
// 4. io.Copy 函数读取响应的内容，然后通过写入 ioutil.Discard 输出流
// 进行丢弃。Copy 返回字节数以及出现的任何错误。
// 5. 每一个结果返回时，fetch 发送一行汇总信息到通道 ch。main 中的第
// 二轮循环接收并且输出那些汇总行。
// 6. 当一个 goroutine 试图在一个通道上进行发送或接收操作时，它会阻
// 塞，直到另一个 goroutine 试图进行接收或发送操作才传递值，并开始处
// 理两个 goroutine。
// 7. 本例中，每一个 fetch 在通道 ch 上发送一个值(ch <- expression)，
// main 函数接收它们(<-ch)。由 main 来处理所有的输出确保每个 goroutine
// 作为一个整体单元处理，这样就避免了两个 goroutine 同时完成造成输出
// 交织所带来的风险。
// ==================================================
// fetchall 和 fetch 一样获取 URL 的内容，但是它并发获取很多 URL 内
// 容，于是这个进程使用的时间不超过耗时最长时间的获取任务，而不是所
// 有获取任务总的时间。
// fetchall 并发获取 URL 并报告它们的时间和大小
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 使用 make 创建一个字符串通道
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 启动一个 goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道 ch 接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道 ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄露资源
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/*
>>> Execution Result:
$ ./fetchall http://baidu.com https://www.baidu.com
0.06s      227  https://www.baidu.com
0.09s       81  http://baidu.com
0.09s elapsed

$ ./fetchall https://golang.org http://gopl.io https://godoc.org
2.75s     4154  http://gopl.io
10.44s    11367  https://godoc.org
Get https://golang.org: dial tcp 216.239.37.1:443: i/o timeout
30.00s elapsed

使用代理后
$ ./fetchall https://golang.org http://gopl.io https://godoc.org
1.40s     4154  http://gopl.io
1.60s    11367  https://godoc.org
Get https://golang.org: dial tcp 216.239.37.1:443: i/o timeout
30.00s elapsed
*/
