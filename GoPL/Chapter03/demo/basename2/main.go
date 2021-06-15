// 简化版利用库函数 string.LastIndex
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}

func basename(s string) string {
	// 如果没找到"/"，则 slash 取值 -1
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

/*
>>> Execution Result:
c
c.d
abc
*/
