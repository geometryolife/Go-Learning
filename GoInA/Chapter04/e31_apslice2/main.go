// 使用append同时增加切片的长度和容量
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}
	fmt.Printf("slice: %d\n", slice)
	fmt.Printf("Length of slice:   %d\n", len(slice))
	fmt.Printf("Capacity of slice: %d\n", cap(slice))

	// 向切片追加一个新元素，将新元素赋值为50
	newSlice := append(slice, 50)
	fmt.Println()
	fmt.Printf("newSlice: %d\n", newSlice)
	fmt.Printf("Length of slice:      %d\n", len(slice))
	fmt.Printf("Capacity of slice:    %d\n", cap(slice))
	fmt.Printf("Length of newSlice:   %d\n", len(newSlice))
	fmt.Printf("Capacity of newSlice: %d\n", cap(newSlice))
}

/*
append操作结束后，newSlice拥有一个全新的底层数组，这个数组的容量是原来的两倍
>>> Execution Result:
slice: [10 20 30 40]
Length of slice:   4
Capacity of slice: 4

newSlice: [10 20 30 40 50]
Length of slice:      4
Capacity of slice:    4
Length of newSlice:   5
Capacity of newSlice: 8
*/
