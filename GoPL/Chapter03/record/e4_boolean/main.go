// 1. bool 型的值或布尔值(boolean)只有两种可能: 真(true)和假(false)。
// 2. if 和 for 语句里的条件，比较操作符(如 == 和 <)，结果均为布尔值。
// 3. 一元操作符(!)表示逻辑取反。
// 4. 考虑到代码风格，布尔表达式 x == true 相对冗长，故化简为 x。
// 5. 布尔值可以由运算符 &&(AND) 以及 ||(OR) 组合运算，这可能引起短
// 路行为: 如果运算符左边的操作数已经能直接确定总体结果，则右边的操
// 作数不会计算在内。
// 6. 布尔值无法隐式转换成数值(0或1)，反之也不行。
// 7. 如果转换操作常常用到，就值得专门为此写个函数: btoi
// 8. 反向转换操作过于简单，无需专门撰写函数，但为了与 btoi 对应，故
// 给出 itob 函数示例代码。
package main

import "fmt"

func main() {
	shortCircuit()
	fmt.Println()
	explicit()
	fmt.Println()
	test()
}

func shortCircuit() {
	var s string

	// 下面的表达式是安全的
	bo := s != "" && s[0] == 'x'
	fmt.Println(bo)

	// 逻辑取反，非真即假
	fmt.Println(!true == false)
	fmt.Println((!true == false) == true)

	c := 'a'

	// && 较 || 优先级更高，如下形式无需加圆括号
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		// ...ASCII 字母或数字
	}
}

func explicit() {
	b := true

	// 如下情况必须使用显示 if
	i := 0
	if b {
		i = 1
		fmt.Println(i)
	}

	// 如第8点所述，获取布尔值十分容易
	j := 1
	fmt.Println(j != 0)
}

func test() {
	x1, x2 := false, true
	fmt.Println(btoi(x1), btoi(x2))

	y1 := 0
	y2 := 1314
	fmt.Println(itob(y1), itob(y2))
}

// 如果 b 为真，btoi 返回1；如果 b 为假，则返回 0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob 报告 i 是否为非零值
func itob(i int) bool { return i != 0 }

/*
>>> Execution Result:
false
true
true

1
true

0 1
false true
*/
