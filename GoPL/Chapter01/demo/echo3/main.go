// Echo3 prints its command-lind arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// 如果不关心格式，只想看值或者调试，那么直接使用Println格式化结果
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0:])
	fmt.Println(os.Args[0])
}
