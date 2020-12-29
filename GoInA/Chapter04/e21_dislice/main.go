// 声明数组和声明切片的不同
package main

import "fmt"

func main() {
	// 创建有3个元素的整型数组
	array := [3]int{10, 20, 30}

	// 创建长度和容量都是3的整型切片
	slice := []int{10, 20, 30}

	fmt.Println("array:")
	for i, v := range array {
		fmt.Printf("%d, %d, %v\n", i, v, &v)
	}
	fmt.Println("slice:")
	for i, v := range slice {
		fmt.Printf("%d, %d, %v\n", i, v, &v)
	}
}

/*
区别: 方括号是否指定数字
>>> Execution Result:
array:
0, 10, 0xc0000140d8
1, 20, 0xc0000140d8
2, 30, 0xc0000140d8
slice:
0, 10, 0xc000014108
1, 20, 0xc000014108
2, 30, 0xc000014108
*/
