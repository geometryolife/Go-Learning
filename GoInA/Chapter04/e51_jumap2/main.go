// 从映射获取值，并通过该值判断键是否存在
package main

import "fmt"

func main() {
	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}

	// 将Blue的代码加入到映射
	colors["Blue"] = "#0000FF"

	// 获取键Blue对应的值
	value := colors["Blue"]

	// 这个键存在吗?
	// 存在
	if value != "" {
		fmt.Println(value)
	}
}

/*
这个例子只返回键对应的值，然后通过判断这个值是不是零值来确定键是否存在。
>>> Execution Result:
#0000FF
*/
