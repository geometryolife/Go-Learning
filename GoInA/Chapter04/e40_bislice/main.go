// 使用空白标识符(下划线)来忽略索引值
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}

	// 迭代每个元素，并显示其值
	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}
}

/*
>>> Execution Result:
Value: 10
Value: 20
Value: 30
Value: 40
*/
