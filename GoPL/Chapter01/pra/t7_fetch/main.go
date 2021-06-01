// 练习1.7: 函数 io.Copy(dst, src) 从 src 读，并且写入 dst。使用它
// 代替 ioutil.ReadAll 来复制响应内容到 os.Stdout，这样不需要装下整
// 个响应数据流的缓冲区。确保检查 io.Copy 返回的错误结果。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		nu, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println(nu)
	}
}
