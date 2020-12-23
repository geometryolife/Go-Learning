// 让Go自动计算声明数组的长度
package main

import "fmt"

func main() {
	// 声明一个整型数组，用具体值初始化每个元素，容量由初始化值决定
	array := [...]int{10, 20, 30, 40, 50}

	for index, value := range array {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}
	fmt.Printf("Length: %d\n", len(array))
}

/*
>>> Execution Result:
Index: 0  Value: 10
Index: 1  Value: 20
Index: 2  Value: 30
Index: 3  Value: 40
Index: 4  Value: 50
Length: 5
*/
