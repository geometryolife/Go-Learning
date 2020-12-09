// 练习2: 修改echo程序，输出参数的索引值，每行一个。
package main

import (
	"fmt"
	"os"
)

func main() {
	// var temp int
	s, sep := "", ""
	for temp, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(temp)
		fmt.Println(arg)
	}
	fmt.Println(s)
}
