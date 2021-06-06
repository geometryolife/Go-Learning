package main

import "fmt"

func main() {
	overflow()
	fmt.Println()
	bitwiseOperator()
	fmt.Println()
	reverse()
	fmt.Println()
	conversion()
	fmt.Println()
	base()
	fmt.Println()
	runeLiteral()
}

func overflow() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1
}

func bitwiseOperator() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)    // 00100010，集合{1, 5}
	fmt.Printf("%08b\n", y)    // 00000110，集合{1, 2}
	fmt.Printf("%08b\n", x&y)  // 00000010，交集{1}
	fmt.Printf("%08b\n", x|y)  // 00100110，并集{1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // 00100100，对称差{2, 5}
	fmt.Printf("%08b\n", x&^y) // 00100000，差集{5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // 元素判定
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // 01000100，集合{2, 6}
	fmt.Printf("%08b\n", x>>1) // 00010001，集合{0, 4}
}

func reverse() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}

func conversion() {
	var apples int32 = 1
	var oranges int16 = 2
	// 非法操作: apples + oranges (int32 与 int16 类型不匹配)
	// var compote int = apples + oranges // 编译错误

	var compote = int(apples) + int(oranges)
	fmt.Println(apples, oranges, compote)

	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	ff := 1e100   // a float64
	ii := int(ff) // 结果依赖实现
	fmt.Println(ff, ii)
}

func base() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 438 666 0666
	x := int64(0xdeadbeef)
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

func runeLiteral() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}

/*
>>> Execution Result:
255 0 1
127 -128 1

00100010
00000110
00000010
00100110
00100100
00100000
1
5
01000100
00010001

bronze
silver
gold

1 2 3
3.141 3
1
1e+100 -9223372036854775808

438 666 0666
3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

97 a 'a'
22269 国 '国'
10 '\n'
*/
