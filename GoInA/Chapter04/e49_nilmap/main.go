// 对nil映射赋值时的语言运行时错误
package main

func main() {
	// 通过声明映射创建一个nil映射
	var colors map[string]string

	// 将Red的代码加入到映射
	colors["Red"] = "#da1337"
}

/*
通过声明一个未初始化的映射来创建一个值为nil的映射(称为nil映射)。nil映射不能
用于存储键值对，否则，会产生一个语言运行时错误。
>>> Execution Result:
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
        /home/Joe/go/src/Go_Learning/GoInA/Chapter04/e49_nilmap/main.go:9 +0x4b
exit status 2
*/
