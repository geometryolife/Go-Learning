// 创建nil切片
/*
>>> File Link:
../e23_nuslice/main.go
*/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 创建nil切片，只需在声明时不做任何初始化
	// var slice []int

	var s1 []int // nil 切片
	s2 := make([]int, 0)

	// 打印 nil 切片和空切片指向的底层数组的地址
	fmt.Printf("s1 pointer: %+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("s2 pointer: %+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))

	// nil 切片和空切片指向的底层数组地址不同
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s1)) == *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
}

/*
结论:
nil 切片和空切片指向的底层数组地址不同，nil 切片不指向任何的实际地址，
空切片指向某个底层数组的地址，无论创建多少个空切片，它们都指向同一个
底层数组的地址，参见下一个例子。nil 切片和空切片的指针指向的底层数组
都没有分配任何存储空间。
>>> Execution Result:
s1 pointer: {Data:0 Len:0 Cap:0}
s2 pointer: {Data:824634150632 Len:0 Cap:0}
false
*/
