// 声明多维切片
package main

import "fmt"

func main() {
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100, 200}}

	fmt.Printf("%d\n", slice)
	fmt.Printf("Length of slice:    %d\n", len(slice))
	fmt.Printf("Length of slice[0]: %d\n", len(slice[0]))
	fmt.Printf("Length of slice[1]: %d\n", len(slice[1]))
}

/*
>>> Execution Result:
[[10] [100 200]]
Length of slice:    2
Length of slice[0]: 1
Length of slice[1]: 2
*/
