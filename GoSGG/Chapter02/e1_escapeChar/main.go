package main

import "fmt"

func main() {
	// 演示转义字符的使用
	fmt.Println("Joe\tTom")
	fmt.Println("hello\nworld")
	fmt.Println("C:\\Users\\Joe\\Desktop")
	fmt.Println("天龙八部雪山飞狐\r张飞")

	// 使用一句输出语句
	fmt.Println("姓名\t年龄\t籍贯\t住址\nJoe\t24\t广西\t广东")
}

/*
>>> Execution Result:
Joe     Tom
hello
world
C:\Users\Joe\Desktop
张飞八部雪山飞狐
姓名    年龄    籍贯    住址
Joe     24      广西    广东
*/
