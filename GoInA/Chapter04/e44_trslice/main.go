// 函数间传递切片
package main

import "fmt"

func main() {
	// 分配包含100万个整型值的切片
	slice := make([]int, 1e6)

	// 将slice传递到函数foo
	slice = foo(slice)
}

// 函数foo接收一个整型切片，并返回这个切片
func foo(slice []int) []int {
	slice2 := slice[2:30:50]

	fmt.Printf("%d\n", slice2)
	fmt.Printf("Length of slice2:   %d\n", len(slice2))
	fmt.Printf("Capacity of slice2: %d\n", cap(slice2))

	return slice2
}

/*
>>> Execution Result:
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Length of slice2:   28
Capacity of slice2: 48
*/
