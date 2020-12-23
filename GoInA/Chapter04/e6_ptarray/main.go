// 访问指针数组的元素
package main

import "fmt"

func main() {
	// 声明包含5个元素的指向整数的数组
	// 用整型指针初始化索引为0和1的数组元素
	array := [5]*int{0: new(int), 1: new(int)}

	// 为索引为0和1的元素赋值
	*array[0] = 10
	*array[1] = 20
	for index, value := range array {
		// fmt.Printf("Index: %d, Value: %d\n", index, *value)
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

/*
fmt.Printf("Index: %d, Value: %d\n", index, *value)
上面这句，输出如下运行时错误报错信息:
Index: 0, Value: 10
Index: 1, Value: 20
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48e87d]
>>> Execution Result:
Index: 0, Value: 824633811152
Index: 1, Value: 824633811160
Index: 2, Value: 0
Index: 3, Value: 0
Index: 4, Value: 0
*/

/*
内存区段错误 Segmentation fault (重定向自SIGSEGV)
内存区段错误（英语：Segmentation fault，经常被缩写为 segfault），又译为记
忆段错误，也称为总线错误（bus error），或访问权限冲突（access violation），
是一种程序错误。它会出现在当程序企图访问CPU无法寻址的内存区段时。当错误发
生时，硬件会通知操作系统，产生了内存访问权限冲突的状况。操作系统通常会产
生核心转储（core dump），以方便程序员进行除错。通常该错误是由调用一个地址，
而该地址为空(NULL)造成的。如链表中调用一个未分配地址的空链表单元的元素。
数组访问越界也可能产生这个错误。
*/
