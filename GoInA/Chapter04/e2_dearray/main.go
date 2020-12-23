// 使用数组字面量声明数组
package main

import "fmt"

func main() {
	// 声明一个包含5个元素的整型数组，用具体值初始化每个元素
	array := [5]int{10, 20, 30, 40, 50}

	for index, value := range array {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}
}

/*
>>> Execution Result:
Index: 0  Value: 10
Index: 1  Value: 20
Index: 2  Value: 30
Index: 3  Value: 40
Index: 4  Value: 50
*/
