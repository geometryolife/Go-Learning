package main

import "fmt"

// 三个名字为 x 的变量，每一个都在不同的词法块中声明；一个在函数体
// 中，一个在 for 语句块中，一个在循环体中，但只有两个块是显式的:
func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (每次迭代一个字母)
	}
}

/*
>>> Execution Result:
HELLO
*/
