// 计算新切片的长度和容量
package main

import "fmt"

func main() {
	// 对底层数组容量是5的切片slice[1:3]来说
	// 长度: 3 - 1 = 2
	// 容量: 5 - 1 = 4
	slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	newSlice := slice[1:3]

	fmt.Printf("Length of slice: %d\n", len(newSlice))
	fmt.Printf("Length of slice: %d\n", cap(newSlice))
	fmt.Printf("%s\n", newSlice)
}

/*
>>> Execution Result:
Length of slice: 2
Length of slice: 4
[Blue Green]
*/
