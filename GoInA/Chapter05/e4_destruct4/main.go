// 使用其它结构类型声明字段
package main

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// admin需要一个user类型作为管理者，并附加权限
type admin struct {
	person user
	level  string
}

func main() {
	// 声明admin类型的变量
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(fred)
}

/*
>>> Execution Result:
{{Lisa lisa@email.com 123 true} super}
*/
