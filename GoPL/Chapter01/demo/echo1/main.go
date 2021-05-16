// echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// 如果不输入参数，则len(os.Args) = 1，此时循环体并未执行
	fmt.Println(len(os.Args))
}
