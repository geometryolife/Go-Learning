package main

import "fmt"

func main() {
	// 日语片假名 "程序"
	s := "プログラム"

	fmt.Printf("%x\n", s)

	// 谓词加空格，以十六进制数形式输出，并在每两个数位间插入空格
	fmt.Printf("% x\n", s)

	r := []rune(s)
	fmt.Printf("%x\n", r)

	// 如果把文字符号类型的切片转换成一个字符串，它会输出各个
	// 文字符号的 UTF-8 编码拼接结果:
	fmt.Println(string(r))

	// 使用 string 函数强制转换数值为字符串会收到警告
	fmt.Println(string(65))
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))

	// 文字符号值非法，被专门的替换字符('\uFFFD')取代
	fmt.Println(string(1234567))

	// 若使用""则打印文字符号值
	// 若使用''则打印文字符号对应的 Unicode 的十进制码点值
	fmt.Println("\uFFFD")
	fmt.Println('\uFFFD')
}

/*
>>> Execution Result:
e38397e383ade382b0e383a9e383a0
e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0
[30d7 30ed 30b0 30e9 30e0]
プログラム
A
A
京
�
�
65533
*/
