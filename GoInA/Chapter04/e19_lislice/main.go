// 通过切片字面量来声明切片
package main

import "fmt"

func main() {
	// 创建字符串切片，其长度和容量都是5个元素
	slice1 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// 创建一个整型切片，其长度和容量都是3个元素
	slice2 := []int{10, 20, 30}

	for i, v := range slice1 {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	fmt.Println()
	for i, v := range slice2 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

/*
>>> Execution Result:
Index: 0, Value: Red
Index: 1, Value: Blue
Index: 2, Value: Green
Index: 3, Value: Yellow
Index: 4, Value: Pink

Index: 0, Value: 10
Index: 1, Value: 20
Index: 2, Value: 30
*/
