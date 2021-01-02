// 使用传统的 for 循环对切片进行迭代
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}

	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}

/*
关键字 range 总是会从切片头部开始迭代，若想对迭代做更多的控制，则可使用
传统的 for 循环。
>>> Execution Result:
Index: 2 Value: 30
Index: 3 Value: 40
*/
