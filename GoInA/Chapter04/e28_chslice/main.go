// 修改切片内容可能导致的结果
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是5个元素
	slice := []int{10, 20, 30, 40, 50}
	fmt.Printf("slice: %d\n", slice)

	// 创建一个新切片，其长度是2个元素，容量是4个元素
	newSlice := slice[1:3]
	fmt.Printf("newSlice: %d\n", newSlice)

	// 修改newSlice索引为1的元素，同时也修改了原来的slice的索引为2的元素
	newSlice[1] = 35
	fmt.Println("After modification:")
	fmt.Printf("slice: %d\n", slice)
	fmt.Printf("newSlice: %d\n", newSlice)
}

/*
>>> Execution Result:
slice: [10 20 30 40 50]
newSlice: [20 30]
After modification:
slice: [10 20 35 40 50]
newSlice: [20 35]
*/
