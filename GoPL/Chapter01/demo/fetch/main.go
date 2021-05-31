// 程序 fetch 展示从互联网获取信息的最小需求，它获取每个指定 URL 的
// 内容，然后不加解析地输出。fetch 来自 curl 这个非常重要的工具。
// fetch 输出从 URL 获取的内容
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// Get 发出一个 GET 请求给指定的 URL
		resp, err := http.Get(url)
		if err != nil {
			// Fprintf 根据格式标识符格式化并写给 w
			// 它返回所写字节的数量和遇到的任何错误
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// Exit 会导致当前程序使用给定的状态代码退出
			// 通常，代码零表示成功，非零表示错误。
			// 程序立即终止：延迟功能不运行。
			// 对于可移植性，状态代码应在[0, 125]范围内。
			os.Exit(1)
		}
		// ReadAll 从 r 读取，直到出现错误或 EOF 并返回其读取的数据。
		// 成功的调用返回 err == nil，而不是 err == EOF
		// 因为 ReadAll 被定义为从 src 读取直到 EOF，所以它不会将
		// 读取的 EOF 视为要报告的错误。
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
