// 练习1.9: 修改 fetch 来输出 HTTP 的状态码，可以在 resp.Status 中找到它。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "http://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	// 输出 HTTP 请求的状态码和状态
	fmt.Printf("Status: %s\n\n", resp.Status)
	fmt.Printf("%s", b)
}
