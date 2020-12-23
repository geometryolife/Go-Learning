// 编译器会阻止类型不同的数组互相赋值
package main

import "fmt"

func main() {
	// 声明第一个包含4个元素的字符串数组
	var array1 [4]string

	// 声明第二个包含5个元素的字符串数组，用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// 把array2的值复制到array1
	// array1 = array2

	// 解决方式
	for i := 0; i < 4; i++ {
		array1[i] = array2[i]
	}

	for index, value := range array1 {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}

/*
数组变量类型包括数组长度和每个元素的类型
>>> Execution Result:
# Go_Learning/GoInA/Chapter04/e8_starray
./main.go:14:9: cannot use array2 (type [5]string) as type [4]string in assignment
解决方案:
Index: 0, Value: Red
Index: 1, Value: Blue
Index: 2, Value: Green
Index: 3, Value: Yellow
*/
