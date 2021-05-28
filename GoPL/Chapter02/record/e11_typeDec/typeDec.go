// 1. 变量或表达式的类型定义这些值应有的特性，例如大小(多少位或多
// 少个元素等)，在内部如何表达，可以对其进行何种操作以及它们所关联
// 的方法。
// 2. type 声明定义一个新的命名类型，它和某个已有类型使用同样的底
// 层类型。命名类型提供了一种方式来区分底层类型的不同用，途甚至可
// 能是不兼容的用途，以便不会无意间将它们混合使用(混用):
// type name underlying-type
// 3. 类型的声明通常出现在包级别，这里命名的类型在整个包中可见，如
// 果名字是可导出的(开头使用大写字母)，其他的包也可以访问它。
// 4. 命名类型的底层类型决定了它的结构和表达方式，以及它支持的
// 内部操作集合，这些内部操作与直接使用底层类型的情况相同。
// 5. 通过 == 和 < 之类的比较操作符，命名类型的值可以与其相同类型
// 的值或者底层类型相同的未命名类型的值相比较。但是不同命名类型的
// 值不能直接比较。
// 6. 命名类型提供了概念上的便利，避免一遍遍地重复写复杂的类型。当
// 底层类型是像 float64 这样简单的类型时，好处就不大了，但是对于我
// 们将讨论到的复杂结构体类型，好处就很大，
/*
>>> File Link:
../../demo/tempconv0/main.go
*/
package main

import (
	tempconv "Go_Learning/GoPL/Chapter02/demo/tempconv0"
	"fmt"
)

func main() {
	fmt.Printf("%g °C\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g °F\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC) // 类型错误: 类型不匹配

	// 如第5点所述
	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f) // 编译错误: 类型不匹配
	// 没有改变参数的值，只改变其类型
	fmt.Println(c == tempconv.Celsius(f)) // "true"

	cc := tempconv.FToC(212.0)
	fmt.Println(cc.String())
	fmt.Printf("%v\n", cc) // 不需要显式调用字符串
	fmt.Printf("%s\n", cc)
	fmt.Println(cc)
	fmt.Printf("%g\n", cc)   // 不调用字符串
	fmt.Println(float64(cc)) // 不调用字符串
}

/*
>>> Execution Result:
100 °C
180 °F
true
true
true
100°C
100°C
100°C
100°C
100
100
*/

// issue:
// cannot define new methods on non-local type tempconv.Celsius
// 无法在非本地类型 tempconv.Celsius 上定义新方法，证明接收者类
// 型无法跨包，故重新定义一个命名类型 Celsius2
// type Celsius2 tempconv.Celsius
// func (c Celsius2) String() string { return fmt.Sprintf("%g°C", c) }
