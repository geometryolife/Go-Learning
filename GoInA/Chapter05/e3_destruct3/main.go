// 不使用字段名，创建结构类型的值
package main

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {
	// 声明 user 类型的变量
	// 每个值也可以分别占一行，不过习惯上这种形式会写在一行里，结尾不需要逗号。
	lisa1 := user{"Lisa", "lisa@email.com", 123, true}
	// 如果不写在一行，结尾需要逗号
	lisa2 := user{
		"Lisa",
		"lisa@email.com",
		123,
		true,
	}

	fmt.Println(lisa1)
	fmt.Println(lisa2)
}

/*
当声明变量时，这个变量的值总是会被初始化。这个值要么用指定的值初始化，要么
用零值(即变量类型的默认值)做初始化。
结构里的每个字段都会用指定的值初始化。
>>> Execution Result:
{Lisa lisa@email.com 123 true}
{Lisa lisa@email.com 123 true}
*/
