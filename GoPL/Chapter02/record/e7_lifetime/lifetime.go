// 1. 生命周期: 指在程序执行过程中变量存在的时间段。
// 2. 包级别变量的生命周期是整个程序的执行时间。
// 3. 局部变量有一个动态的生命周期:
// 每次执行声明语句时创建一个新的实体，变量一直生存到它变得不可访问，
// 这时它占用的存储空间被回收。
// 4. 函数的参数和返回值也是局部变量，它们在其闭包函数被调用时创建。
// 5. 垃圾回收器如何知道一个变量是否该被回收? 基本思想:
// 每一个包级别的变量，以及每一个当前执行函数的局部变量，可以作为追
// 溯该变量的路径的源头，通过指针和其他方式的引用可以找到变量。如果
// 变量路径不存在，那么变量变得不可访问，因此它不会影响任何其他的计
// 算过程。
// 6. 因为变量的生命周期是通过它是否可达来确定的，所以局部变量可在
// 包含它的循环的一次迭代之外继续存活。即使包含它的循环已经返回，它
// 的存在还可能延续。
// 7. 编译器可以选择在堆上或栈上分配局部变量，但令人惊讶的是，这个
// 选择不是由使用 var 还是 new 关键字声明变量来确定的。
// 8. 任何情况下，逃逸的概念使你不需要额外费心来写正确的代码，但要
// 记住它在性能优化的时候是有好处的，因为每一次变量逃逸都需要一次
// 额外的内存分配过程。
// 9. Golang的垃圾回收机制对写出正确的程序有巨大的帮助，但是避免不
// 了考虑内存的负担。不需要显式分配和释放内存，但是变量的生命周期是
// 写出高效程序所必须清楚的。
// 10. 在长生命周期对象中保持短生命周期对象不必要的指针，特别是在全
// 局变量中，会阻止垃圾回收器回收短生命周期的对象空间。
/*
>>> File Link:
../../../Chapter01/demo/lissajous/main.go
*/
package main

import "fmt"

var global *int

func main() {
	f()
	fmt.Println(*global)

	g()
	// fmt.Println(*y)
}

func f() {
	// x 一定使用堆空间，因为它在f函数返回以后还可以从
	// global变量访问，尽管它被声明为一个局部变量
	// 这种情况我们说 x 从 f() 中逃逸
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	// 当 g() 函数返回时，变量 *y 变得不可访问，可回收
	// 因为 *y 没有从 g() 中逃逸，所以编译器可以安全地在栈上
	// 分配 *y，即便使用 new 函数创建它。
	*y = 1
}

/*
>>> Execution Result:
1
*/