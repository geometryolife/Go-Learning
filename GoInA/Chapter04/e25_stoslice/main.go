// 使用切片创建切片
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是5个元素
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println("slice:")
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d, Address: %v\n", i, v, &slice[i])
	}

	// 创建一个新切片，其长度为2个元素，容量为4个元素
	newSlice := slice[1:3]

	fmt.Println("newSlice:")
	for i, v := range newSlice {
		fmt.Printf("Index: %d, Value: %d, Address: %v\n", i, v, &newSlice[i])
	}
}

/*
通过地址可以知道，这两个切片共享同一个底层数组
>>> Execution Result:
slice:
Index: 0, Value: 10, Address: 0xc000014210
Index: 1, Value: 20, Address: 0xc000014218
Index: 2, Value: 30, Address: 0xc000014220
Index: 3, Value: 40, Address: 0xc000014228
Index: 4, Value: 50, Address: 0xc000014230
newSlice:
Index: 0, Value: 20, Address: 0xc000014218
Index: 1, Value: 30, Address: 0xc000014220
*/
