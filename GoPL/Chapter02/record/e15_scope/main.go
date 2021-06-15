package main

import "fmt"

func f() {}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f) // "f"; 局部变量 f 覆盖了包级函数 f
	fmt.Println(g) // "g"; 包级变量
	// fmt.Println(h) // 编译错误: 为定义 h
}

/*
>>> Execution Result:
f
g
*/
