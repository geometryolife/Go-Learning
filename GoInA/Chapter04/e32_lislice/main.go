// 使用切片字面量声明一个字符串切片
package main

import "fmt"

func main() {
	// 创建字符串切片，其长度和容量都是5个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 创建切片时的第三个索引可以限制容量
	slice := source[2:3:4]
	fmt.Printf("Capacity of source: %d\n", cap(source))
	fmt.Printf("Capacity of slice:  %d\n", cap(slice))
}

/*
>>> Execution Result:
Capacity of source: 5
Capacity of slice:  2
*/
