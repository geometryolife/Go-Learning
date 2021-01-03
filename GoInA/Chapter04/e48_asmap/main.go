// 为映射赋值
package main

import "fmt"

func main() {
	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}

	// 将Red的代码加入到映射
	colors["Red"] = "#da1337"

	fmt.Printf("%s", colors)
}

/*
>>> Execution Result:
map[Red:#da1337]
*/
