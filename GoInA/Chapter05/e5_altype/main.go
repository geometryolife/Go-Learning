// 基于int64声明一个新类型
package main

import "fmt"

// Duration for test.
type Duration int64

func main() {
	var dur Duration

	// 给不同类型的变量赋值会产生编译错误
	// dur = int64(1000)

	// 修正:
	dur = Duration(1000)

	fmt.Println(dur)
}

/*
两种不同类型的值即便互相兼容，也不能互相赋值。编译器不会对不同类型的值做隐式转换。
>>> Execution Result:
# Go_Learning/GoInA/Chapter05/e5_altype
./main.go:10:6: cannot use int64(1000) (type int64) as type Duration in assignment
修正输出:
1000
*/
