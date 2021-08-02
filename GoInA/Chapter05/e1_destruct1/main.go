// 声明一个结构类型
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
	// 使用结构类型声明变量，并初始化为其零值
	// 声明 user 类型的变量
	var bill user

	fmt.Println(bill)
}

/*
当声明变量时，这个变量的值总是会被初始化。这个值要么用指定的值初始化，要么
用零值(即变量类型的默认值)做初始化。
结构里的每个字段都会用零值初始化。
>>> Execution Result:
{  0 false}
*/
