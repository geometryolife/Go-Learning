// 使用结构字面量来声明一个结构类型的变量
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
	// 声明user类型的变量，并初始化所有字段
	// 在不同行声明每个字段的名字以及对应的值，字段名与值用冒号分隔，
	// 每一行以逗号结尾，此形式对字段的声明顺序没有要求。
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	fmt.Println(lisa)
}

/*
当声明变量时，这个变量的值总是会被初始化。这个值要么用指定的值初始化，要么
用零值(即变量类型的默认值)做初始化。
结构里的每个字段都会用指定的值初始化。
>>> Execution Result:
{Lisa lisa@email.com 123 true}
*/
