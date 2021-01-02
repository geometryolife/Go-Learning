// 显示索引越界的运行时错误
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是5个元素
	slice := []int{10, 20, 30, 40, 50}

	// 创建一个新切片，其长度为2个元素，容量为4个元素
	newSlice := slice[1:3]

	// 修改newSlice索引为3的元素，这个元素对于newSlice来说并不存在
	newSlice[3] = 45
	// 注释上一行代码，再编译运行，可看到newSlice只有索引0-1两个元素
	fmt.Printf("%d\n", newSlice)
}

/*
>>> Execution Result:
panic: runtime error: index out of range [3] with length 2

goroutine 1 [running]:
main.main()
        /home/Joe/go/src/Go_Learning/GoInA/Chapter04/e29_ufslice/main.go:14 +0x61
exit status 2
*/
