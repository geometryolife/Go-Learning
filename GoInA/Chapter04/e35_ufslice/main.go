// 设置容量大于已有容量的语言运行时错误
package main

import "fmt"

func main() {
	// 创建字符串切片，其长度和容量都是5个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 这个切片操作试图设置容量为4，这比可用的容量大
	slice := source[2:3:6]
	fmt.Printf("Length of slice: %d\n", len(slice))
	fmt.Printf("Capacity of slice: %d\n", cap(slice))
	fmt.Printf("%s\n", slice)
}

/*
>>> Execution Result:
panic: runtime error: slice bounds out of range [::6] with capacity 5

goroutine 1 [running]:
main.main()
        /home/Joe/go/src/Go_Learning/GoInA/Chapter04/e35_ufslice/main.go:11 +0x5b
exit status 2
*/
