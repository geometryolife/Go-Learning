// 使用长度和容量声明整型切片
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度为3个元素，容量为5个元素
	slice := make([]int, 3, 5)

	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

/*
初始化后并不能访问所有的数组元素
>>> Execution Result:
Index: 0, Value: 0
Index: 1, Value: 0
Index: 2, Value: 0
*/
