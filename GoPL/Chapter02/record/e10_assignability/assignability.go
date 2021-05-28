// 1. 赋值语句是显式形式的赋值，但是程序中很多地方的赋值是隐式的:
// (1)一个函数调用隐式地将参数的值赋给对应参数的变量；
// (2)一个return语句隐式地将return操作数赋值给结果变量；
// (3)复合类型的字面量表达式，例如slice。
// 2. 不论隐式还是显式赋值，如果左边的(变量)和右边的(值)类型相同,
// 它就是合法的。通俗地说，赋值只有在值对于变量类型是可赋值的时候
// 才合法。
// 3. 可赋值性根据类型的不同有着不同的规则，在引入新类型的时候解
// 释相应的规则。
// 4. nil 可以被赋给任何接口变量或引用类型，常量有更灵活的可赋值
// 性规则来规避显式的转换。
// 5. 两个值是否可以使用 == 和 != 进行比较与可赋值性相关:
// 可比较性 -> 任何比较中，第一个操作数必须是可分配给第二个操作数
// 的类型，反之亦然。
package main

import "fmt"

func main() {
	sli1()
	sli2()
}

// 隐式赋值
func sli1() {
	medals := []string{"gold", "silver", "bronze"}
	for _, medal := range medals {
		fmt.Println(medal)
	}
}

// 显式赋值
func sli2() {
	var medals []string

	// 为声明的空切片追加元素
	for i := 0; i < 3; i++ {
		medals = append(medals, "")
		// 查看切片容量的变化情况
		fmt.Println(cap(medals))
	}

	medals[0] = "gold"
	medals[1] = "silver"
	medals[2] = "bronze"
	for _, medal := range medals {
		fmt.Println(medal)
	}
}

/*
>>> Execution Result:
gold
silver
bronze
1
2
4
gold
silver
bronze
*/
