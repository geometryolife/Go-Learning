// 最简单的 Web 服务器
package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
}
