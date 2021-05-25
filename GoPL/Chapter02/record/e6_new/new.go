// 1. 另一种创建变量的方式: 使用内置的new()函数。
// 2. new(T)创建一个未命名的T类型变量，初始化为T类型的零值，
// 并返回其地址(地址类型为*T)。
// 3. 使用new创建的变量和取其地址的普通局部变量没有什么不同，
// 只是不需要引入(和声明)一个虚拟名字，通过new(T)就可以直接在
// 表达式中使用。因此new只是语法上的便利，不是一个基本概念。
// 4. new是一个预声明的函数，不是一个关键字，所以它可以重
// 定义为另外的其他类型:
// func delta(old, new int) int { return new - old }
package main

import "fmt"

func main() {
	p := new(int)   // *int类型的p，指向未命名的int变量
	fmt.Println(*p) // 输出 "0"
	*p = 2          // 把未命名的int设置为2
	fmt.Println(*p) // 输出 "2"

	// 每一次调用new()，返回一个具有唯一地址的不同变量
	p1 := new(int)
	p2 := new(int)
	fmt.Println(p1 == p2) // "false"

	// 书中说new返回唯一地址有一个例外:
	// 两个变量类型不携带任何信息且是零值，有相同的地址，
	// 例如struct{}或[0]int
	// 经过实验，在新版本的Golang中已经统一了唯一性
	// go version: go1.13.8 linux/amd64
	q1 := new(struct{})
	q2 := new(struct{})
	fmt.Println(q1 == q2) // "false"

	w1 := new([0]int)
	w2 := new([0]int)
	fmt.Println(w1 == w2) // "false"

	// newInt1() & newInt2()有同样的行为
	fmt.Println(newInt1())
	fmt.Println(newInt2())
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

/*
>>> Execution Result:
0
2
false
false
false
0xc0000140e0
0xc0000140e8
*/
