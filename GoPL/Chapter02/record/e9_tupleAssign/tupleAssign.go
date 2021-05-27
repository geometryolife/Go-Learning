// 1. 多重赋值: 允许多个变量一次性被赋值。
// 2. 在实际更新变量前，右边所有的表达式被推演，当变量同时出现在
// 赋值符两侧的时候这种形式特别有用。
// 3. 多重赋值可以使一个普通的赋值序列变得紧凑:
// i, j, k = 2, 3, 5
// 4. 从风格上考虑，如果表达式比较复杂，则避免使用多重赋值形式，
// 一系列独立的语句更易读。
// 5. 如果map查询、类型断言或者通道接收动作出现在两个结果的赋值语
// 句中，都会产生一个额外的布尔型结果:
// v, ok = m[key] // map 查询
// v, ok = x.(T)  // 类型断言
// v, ok = <-ch   // 通道接收
// 6. 像变量声明一样，可以将不需要的值赋值给空标识符:
// _, err = io.Copy(dst,src) // 丢弃字节个数
// _, ok = x.(T)             // 检查类型但丢弃结果
package main

import "fmt"

func main() {
	change()
	fmt.Println(gcd(128, 256))
	fmt.Println(fib(10))
}

// 交换两个变量
func change() {
	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

	a := [3]int{1, 2, 3}
	fmt.Println(a)
	a[0], a[2] = a[2], a[0]
	fmt.Println(a)
}

// 计算两个整数的最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 计算斐波那契数列的第n个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

/*
>>> Execution Result:
1 2
2 1
[1 2 3]
[3 2 1]
128
55
*/
