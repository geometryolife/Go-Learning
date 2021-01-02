// range 提供了每个元素的副本
package main

import "fmt"

func main() {
	// 创建一个整型切片，其长度和容量都是4个元素
	slice := []int{10, 20, 30, 40}

	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}
}

/*
range 创建了每个元素的副本，而不是直接返回对该元素的引用，如果使用该值
变量的地址作为指向每个元素的指针，就会造成错误。
因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以value的
地址总是相同的。
借助切片变量和索引值，可以获取每个元素的地址值。
>>> Execution Result:
Value: 10 Value-Addr: C0000140D0 ElemAddr: C000016180
Value: 20 Value-Addr: C0000140D0 ElemAddr: C000016188
Value: 30 Value-Addr: C0000140D0 ElemAddr: C000016190
Value: 40 Value-Addr: C0000140D0 ElemAddr: C000016198
*/
