// 把同样类型的一个数组赋值给另外一个数组
package main

import "fmt"

func main() {
	// 声明第一个包含5个元素的字符串数组
	var array1 [5]string

	// 声明第二个包含5个元素的字符串数组，用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// 把array2的值复制到array1
	array1 = array2

	for index, value := range array1 {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}

/*
>>> Execution Result:
Index: 0, Value: Red
Index: 1, Value: Blue
Index: 2, Value: Green
Index: 3, Value: Yellow
Index: 4, Value: Pink
*/
