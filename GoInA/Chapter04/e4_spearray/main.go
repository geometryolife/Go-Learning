// 声明数组并指定特定元素的值
package main

import "fmt"

func main() {
	// 声明一个有5个元素的数组，用具体值初始化索引为1和2的元素
	// 其余元素保持零值
	array := [5]int{1: 10, 2: 20}

	for index, value := range array {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}
}

/*
>>> Execution Result:
Index: 0  Value: 0
Index: 1  Value: 10
Index: 2  Value: 20
Index: 3  Value: 0
Index: 4  Value: 0
*/
