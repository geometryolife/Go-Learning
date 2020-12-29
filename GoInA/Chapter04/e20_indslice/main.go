// 使用索引声明切片
package main

import "fmt"

func main() {
	// 创建字符串切片，使用空字符串初始化第100个元素
	slice := []string{99: ""}

	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
}

/*
使用字面量来创建切片时，可以使用长度和容量作为索引来指定切片的长度和容量
>>> Execution Result:
Index: 0, Value:
Index: 1, Value:
Index: 2, Value:
Index: 3, Value:
Index: 4, Value:
Index: 5, Value:
...
Index: 94, Value:
Index: 95, Value:
Index: 96, Value:
Index: 97, Value:
Index: 98, Value:
Index: 99, Value:
*/
