// 1. 声明变量:
// var name type = expression
// 2. 类型和表达式部分可以省略一个，但不能都省略
// 3. 零值:
// 数字 -> 0
// 布尔型 -> false
// 字符串 -> ""(空字符串)
// 接口、引用类型(slice、指针、map、通道、函数) -> nil
// 复合类型(数组、结构体) -> 其所有元素或成员的零值
// 4. 零值机制的作用: 保证每个变量是良好定义的，同时简化代码
// 5. Go程序员经常花费精力来使复杂类型的零值有意义，以便变量
// 一开始就处于一个可用状态。
// 6. 包级别的初始化在main开始之前进行，局部变量初始化和声明一样
// 在函数执行期间进行。
// 7. 变量可以通过调用返回多个值的函数进行初始化
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string
	var a int
	var b float32
	var c float64
	// 声明变量列表
	var i, j, k int                    // int, int, int
	var bo, fl, st = true, 2.3, "four" // bool, float64, string

	// 输出空字符串，而不是一些错误或不可预料的行为
	fmt.Println(s)
	fmt.Println(a, b, c)
	fmt.Println(i, j, k)
	fmt.Println(bo, fl, st)

	// 调用多返回值函数对变量进行初始化
	var va, ty = multipleValues()
	fmt.Println(va, ty)
}

func multipleValues() (string, reflect.Type) {
	v := "Hello Go!"
	t := reflect.TypeOf(v)

	return v, t
}

/*
>>> Execution Result:

0 0 0
0 0 0
true 2.3 four
Hello Go! string
*/
