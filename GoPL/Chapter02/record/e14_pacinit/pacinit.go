// 1. 包的初始化从初始化包级别的变量开始，这些变量按照声明顺序初始化，
// 在依赖已解析完毕的情况下，根据依赖的顺序进行。(包的初始化首先按照
// 声明的顺序初始化包级别的变量，除了首先解析依赖项)。
// 2. 如果包由多个 .go 文件组成，初始化按照编译器收到的文件的顺序进
// 行: go 工具会在调用编译器前将 .go 文件进行排序。
// 3 对于包级别的每一个变量，生命周期从其值被初始化开始，但是对于其
// 他一些变量，比如数据表，初始化表达式可能不是设置其初始化值的最简
// 单方法。在这种情况下，init 函数机制可能更简单。任何文件都可以包
// 含任意数量的 init() 函数。
// 4. init 函数不能被调用和被引用，另一方面，它也是普通的函数。在每
// 一个文件里，当程序启动的时候，init 函数按照它们声明的顺序自动执行。
// 5. 包的初始化按照在程序中导入的顺序来进行，依赖顺序优先，每次初始
// 化一个包。
// 6. 初始化过程是自下而上的，main 包最后初始化。所以程序的 main 函
// 数开始执行前，所有的包已初始化完毕。
/*
>>> File Link:
../../demo/popcount/popcount.go
*/
package main

import "fmt"

// 如第1点所述:
var a = b + c // 最后把 a 初始化为3
var b = f()   // 通过调用 f 接着把 b 初始化为2
var c = 1     //首先初始化为1

func f() int { return c + 1 }

func main() {
	fmt.Println(a, b, c)
}

/*
>>> Execution Result:
3 2 1
*/