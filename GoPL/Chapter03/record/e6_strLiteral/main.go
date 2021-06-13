// 1. 字符串的值可以直接写成字符串字面量(string literal)，形式上就是带
// 双引号的字节序列。
// 2. Golang 按 UTF-8 编码，在源码中可以将 Unicode 码点写入字符串字面量。
// 3. 转义序列以反斜杠(\)开始，可以将任意值的字节插入字符串中。
// 4. 源码中的字符串可以包含十六进制或八进制的任意字节:
// 十六进制的转义序列写成: \xhh
// 八进制的转义序列写成: \ooo
// hh 代表两位十六进制数字，ooo 代表三位八进制数字(0~7)，且不能超过 \377。
// 5. 原生的字符串字面量的书写形式是 `...`，使用反引号而不是双引号。
// 6. 原生的字符串字面量内，转义序列不起作用；实质内容与字面书写严格一
// 致，包括反斜杠和换行符，因此，在程序源码中，原生的字符串字面量可以展
// 开多行。唯一的特殊处理是回车符会被删除(换行符会保留)，使得同一字符串
// 在所有平台上的值都相同，包括习惯在文本文件存入换行符的系统。
// 7. 正则表达式往往含有大量反斜杠，可以方便地写成原生的字符串字面量。
// 原生的字符串字面量也适用于 HTML 模版、JSON 字面量、命令行提示信息，
// 以及需要多行文本表达的场景。
package main

import "fmt"

func main() {
	const s = "Hello, 世界"
	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
	go command [arguments]
	...`

	fmt.Println(s)
	fmt.Println(GoUsage)
}

/*
>>> Execution Result:
Hello, 世界
Go is a tool for managing Go source code.
        Usage:
        go command [arguments]
        ...
*/
