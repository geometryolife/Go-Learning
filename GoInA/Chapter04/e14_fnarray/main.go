// 使用值传递，在函数间传递大数组
package main

import "fmt"

func main() {
	// 声明一个需要 8 MB 的数组
	var array [1e6]int

	// 将数组传递给函数foo
	foo(array)

}

// 函数foo接受一个100万个整型值的数组
func foo(array [1e6]int) {
	array2 := array

	// 若将条件上限设置为 1e6，则遍历所有地址
	for i := 0; i < 6; i++ {
		fmt.Printf("The address of the %d subscript: %v\n", i, &array2[i])
	}
}

/*
>>> Execution Result:
The address of the 0 subscript: 0xc000080000
The address of the 1 subscript: 0xc000080008
The address of the 2 subscript: 0xc000080010
The address of the 3 subscript: 0xc000080018
The address of the 4 subscript: 0xc000080020
The address of the 5 subscript: 0xc000080028
*/
