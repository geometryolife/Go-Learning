// 使用3个索引创建切片
package main

import "fmt"

func main() {
	// 创建字符串切片，其长度和容量都是5个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 将第三个元素切片，并限制容量，其长度为1个元素，容量为2个元素
	slice := source[2:3:4]
	fmt.Printf("Length of slice:   %d\n", len(slice))
	fmt.Printf("Capacity of slice: %d\n", cap(slice))
	fmt.Printf("%s\n", slice)
}

/*
>>> Execution Result:
Length of slice:   1
Capacity of slice: 2
[Plum]
*/
