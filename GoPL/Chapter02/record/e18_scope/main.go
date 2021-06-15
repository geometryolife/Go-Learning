// 像 for 循环一样，除了本身的主体块之外，if 和 switch 语句还
// 会创建隐式的词法块。
package main

import "fmt"

// if 语句创建隐式的词法块，else if 和 else 嵌套在 if 中，
// 所以 if 语句中初始化声明的变量在嵌套语句中是可见的
func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	// fmt.Println(x, y) // 编译错误: x 与 y 在这里不可见
}

func f() int {
	// return 0
	var a int

	fmt.Println("请输入一个整数:")
	fmt.Scan(&a)

	return a
}

func g(x int) int {
	if x < 2 {
		return x
	} else {
		return x + 10
	}
}

/*
>>> Execution Result:
请输入一个整数:
0
0
请输入一个整数:
1
1 1
请输入一个整数:
2
2 12
*/
