// 如何计算长度和容量
package main

import "fmt"

func main() {
	// 对底层数组容量是k的切片slice[i:j]来说
	// 长度: j - i
	// 容量: k - i
	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3]

	// 长度:
	fmt.Printf("slice length:    %d\n", len(slice))
	fmt.Printf("newSlice length: %d\n", len(newSlice))
	// 容量:
	fmt.Printf("slece capacity:    %d\n", cap(slice))
	fmt.Printf("newSlice capacity: %d\n", cap(newSlice))
}

/*
>>> Execution Result:
slice length:    5
newSlice length: 2
slece capacity:    5
newSlice capacity: 4
*/
