// Golang的源文件后缀为 .go
// 当前文件属于 main 包，在 Go 中每个文件必须属于一个包，无法单独存在
package main

// 表示引入一个包，包名 fmt，引入该包后，可以使用该包的代码
import "fmt"

// func 是关键字，后跟函数名 main，是程序的主函数，即程序的入口
func main() {
	fmt.Println("Hello World!")
}
