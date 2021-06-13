// 1. 字符串是不可变的字节序列，可包含任意数据，包括0值字节，但主要是
// 人类可读的文本。
// 2. 习惯上，文本字符串被解读成按 UTF-8 编码的 Unicode 码点(文字符号)
// 序列。
// 3. 内置的 len 函数返回字符串的字节数(并非文字符号的数目)，下标访问操
// 作s[i]则取得第 i 个字符，其中 0≤i<len(s)。
// 4. 字符串的第 i 个字节不一定就是第 i 个字符，因为非 ASCII 字符的
// UTF-8 码点需要两个字节或多个字节。
// 5. 字串生成操作 s[i:j] 产生一个新字符串，内容取自原字符串的字节，
// 下标从 i (含边界值) 开始，直到 j (不含边界值)。结果的大小是 j-i 个
// 字节。操作数 i 与 j 的默认值分别是 0 (字符串起始位置) 和 len(s)
// (字符串终止位置)，若省略 i 或 j，或两者，则取默认值。
// 6. 字符串可以通过比较运算符做比较，如 == 和 <；比较运算符按字节进行，
// 结果服从本身的字典排序。
// 7. 尽管可以将新值赋予字符串变量，但是字符串值无法改变: 字符串值本身
// 所包含的字节序列永不可变。
// 8. 不可变意味着两个字符串能安全地共用同一段底层内存，使得复制任何长
// 度字符串的开销都低廉。类似地，字符串 s 及其字串(如 s[7:])可以安全地
// 共用数据，因此子串生成操作的开销低廉。这两种情况都没有分配新内存。
package main

import "fmt"

func main() {
	subscript()
	fmt.Println()
	substring()
	fmt.Println()
	app()
}

func subscript() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Printf("%c, %c\n", s[0], s[7])

	// 访问许可范围以外的字节会触发宕机异常
	// c := s[len(s)] // 宕机: 下标越界
	// panic: runtime error: index out of range [12] with length 12
}

func substring() {
	a := "你好世界"
	b := "Hello World"
	s := "hello, world"

	fmt.Println(a[1:2])
	fmt.Println(b[1:2])

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	// 加号(+)运算符连接两个字符串而生成一个新字符串
	fmt.Println("goodbye" + s[5:])

	// 比较
	fmt.Println(s > b)
	fmt.Println(s == b)
}

func app() {
	s := "left foot"
	t := s
	s += ", right foot"

	fmt.Println(s)
	fmt.Println(t)

	// 字符串不可改变，不允许修改字符串内部的数据
	// s[0] = 'L' // 编译错误: s[0] 无法赋值
}

/*
>>> Execution Result:
12
104 119
h, w

�
e
hello
hello
world
hello, world
goodbye, world
true
false

left foot, right foot
left foot
*/
