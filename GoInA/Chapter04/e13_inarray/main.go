// 使用索引为多维数组赋值
package main

import "fmt"

func main() {
	// 声明两个不同的二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int

	// 为每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40

	// 将array2的值复制给array1
	array1 = array2

	// 将array1的索引为1的维度复制到一个同类型的新数组里
	var array3 [2]int = array1[1]

	// 将外层数组的索引为1、内层数组的索引为0的整型值复制到新的整型变量里
	var value int = array1[1][0]

	for i, v := range array3 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	fmt.Printf("The integer value is %d\n", value)
}

/*
>>> Execution Result:
Index: 0, Value: 30
Index: 1, Value: 40
The integer value is 30
*/
