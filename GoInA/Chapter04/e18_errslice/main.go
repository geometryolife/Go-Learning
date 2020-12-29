// 容量小于长度的切片会在编译时报错
package main

import "fmt"

func main() {
	// 创建一个整型切片，使其长度大于容量
	// slice := make([]int, 5, 3)
	// 修改:
	slice := make([]int, 5, 5)

	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

/*
>>> Execution Result:
报错信息:
# Go_Learning/GoInA/Chapter04/e18_errslice
./main.go:8:15: len larger than cap in make([]int)
修改:
Index: 0, Value: 0
Index: 1, Value: 0
Index: 2, Value: 0
Index: 3, Value: 0
Index: 4, Value: 0
*/
