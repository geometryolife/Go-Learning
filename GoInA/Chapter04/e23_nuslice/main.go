// 声明空切片
/*
>>> File Link:
../e22_nilslice/main.go
*/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 使用make创建空的整型切片
	// slice := make([]int, 0)

	// 使用切片字面量创建空的整型切片
	// slice := []int{}

	s1 := make([]int, 0) // 空切片 s1
	s2 := make([]int, 0) // 空切片 s2
	s3 := []int{}
	s4 := []int{}

	fmt.Printf("s1 pointer: %+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("s2 pointer: %+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("s3 pointer: %+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
	fmt.Printf("s4 pointer: %+v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
}

/*
>>> Execution Result:
s1 pointer: {Data:824634150568 Len:0 Cap:0}
s2 pointer: {Data:824634150568 Len:0 Cap:0}
s3 pointer: {Data:824634150568 Len:0 Cap:0}
s4 pointer: {Data:824634150568 Len:0 Cap:0}
*/
