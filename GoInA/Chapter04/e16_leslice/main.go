// 使用长度声明一个字符串切片
package main

import "fmt"

func main() {
	// 创建一个字符串切片，使其长度和容量都是5个元素
	slice := make([]string, 5)

	for index := 0; index < len(slice); index++ {
		fmt.Printf("Index: %d, Value: %s, Address: %v\n", index, slice[index],
			&slice[index])
	}
}

/*
当使用make函数创建长度和容量相等的的字符串切片时，每个元素初始化为空字符串
>>> Execution Result:
Index: 0, Value: , Address: 0xc000078000
Index: 1, Value: , Address: 0xc000078010
Index: 2, Value: , Address: 0xc000078020
Index: 3, Value: , Address: 0xc000078030
Index: 4, Value: , Address: 0xc000078040
*/
