// 使用 for range 迭代切片
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}

	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
}

/*
关键字range返回两个值，第一个是当前迭代到的索引位置，第二个值是
该位置对应元素值的一份副本
>>> Execution Result:
Index: 0 Value: 10
Index: 1 Value: 20
Index: 2 Value: 30
Index: 3 Value: 40
*/
