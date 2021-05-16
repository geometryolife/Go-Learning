// 练习2: 修改echo程序，输出参数的索引和值，每行一个。
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(index, arg)
	}
	fmt.Println(s)
}
