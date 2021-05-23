// 1. 名称包括:
// 函数、变量、常量、类型、语句标签、包
// 2. Golang有25个关键字
// 3. 实体的声明作用域:
// 函数内 -> 函数局部可见
// 函数外 -> 对包内所有源文件可见
// 实体名首字母大写 -> 可导出的，其在包外可见、可访问
// 总结: 实体的第一个字母的大小写决定了其可见性是否跨包
// 4. 包名总是由小写字母组成
// 5. 名称没有长度限制，但Golang编程风格偏向短名称，驼峰式
// 6. 像ASCII、HTML这样的首字母缩写词，通常使用相同大小写:
// htmlEscape、HTMLEscape、escapeHTML
package main

import "fmt"

var htmlEscape = "Hello HTML!"

func main() {
	htmlEscapeLocal := "local HTML."

	name()
	fmt.Println(htmlEscapeLocal)

	htmlEscape = "Hello HTML has been changed!"
	name()
}

func name() {
	fmt.Println(htmlEscape)
}

/*
>>> Execution Result:
Hello HTML!
local HTML.
Hello HTML has been changed!
*/
