// 声明一个数组，并设置为零值
package main

import "fmt"

func main() {
	// 声明一个包含5个元素的整型数组，数组内每个元素都初始化为对应类型的零值
	// 一旦声明，数组里存储的数据类型和数组长度就都不能改变了。
	var array [5]int

	for index, value := range array {
		fmt.Printf("Index: %d  Value: %d\n", index, value)
	}
}

/*
>>> Execution Result:
Index: 0  Value: 0
Index: 1  Value: 0
Index: 2  Value: 0
Index: 3  Value: 0
Index: 4  Value: 0
*/
