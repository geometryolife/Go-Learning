// 1. 赋值语句用来更新变量所指的值
// 2. 每一个算术和二进制位操作符有一个对应的赋值操作符:
// count[x] *= scale -> 避免了在表达式中重复变量本身
// 3. 数字变量也可以通过 ++ 和 -- 语句进行递增和递减
package main

import "fmt"

func main() {
	var x int
	scale := 5
	p := new(bool)
	type Person struct {
		name string
		age  int
	}
	var person Person
	var count [3]int

	fmt.Println(x, *p, person.name, count[x])

	x = 1                       // 有名称的变量
	*p = true                   // 间接变量
	person.name = "bob"         // 结构体成员
	count[x] = count[x] * scale // 数组或slice或map的元素

	fmt.Println(x, *p, person.name, count[x])

	v := 1
	v++ // 等同于 v = v + 1; v 变成 2
	fmt.Println(v)
	v-- // 等同于 v = v - 1; v 变成 1
	fmt.Println(v)
}

/*
>>> Execution Result:
0 false  0
1 true bob 0
2
1
*/
