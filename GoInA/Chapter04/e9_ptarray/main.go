// 把一个指针数组赋值给另一个
package main

import "fmt"

func main() {
	// 声明第一个包含3个元素的指向字符串的指针数组
	var array1 [3]*string

	// 声明第二个包含3个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array2 := [3]*string{new(string), new(string), new(string)}

	// 使用颜色为每个元素赋值
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	// 将array2复制给array1
	array1 = array2

	fmt.Println("array1:")
	for i, v := range array1 {
		fmt.Printf("Index: %d, Address: %v value is %s\n", i, v, *v)
	}
	fmt.Println("array2:")
	for i, v := range array2 {
		fmt.Printf("Index: %d, Address: %v value is %s\n", i, v, *v)
	}

}

/*
>>> Execution Result:
array1:
Index: 0, Address: 0xc000010200 value is Red
Index: 1, Address: 0xc000010210 value is Blue
Index: 2, Address: 0xc000010220 value is Green
array2:
Index: 0, Address: 0xc000010200 value is Red
Index: 1, Address: 0xc000010210 value is Blue
Index: 2, Address: 0xc000010220 value is Green
*/
