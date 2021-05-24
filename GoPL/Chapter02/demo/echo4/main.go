// echo4 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

// Bool()函数使用指定的名称，默认值和用法字符串定义一个bool标志
// 返回值是存储标志值的bool变量的地址。
var n = flag.Bool("n", false, "omit trailing newline")

// 字符串定义具有指定名称，默认值和用法字符串的字符串标志
// 返回值是存储标志值的字符串变量的地址
var sep = flag.String("s", " ", "separator")

func main() {
	// Parse()函数从os.Args[1:]解析命令行标志
	// 必须在定义所有标志之后并且在程序访问标志之前调用
	flag.Parse()

	// Args()函数以字符串切片的方式返回非标志命令行参数
	// 参数1 -> Join()将字符串切片的元素连接起来以创建单个字符串
	// 参数2 -> 分隔符字符串放置在结果字符串中的元素之间
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(flag.Args())
}

/*
>>> Execution Result:
$ ./echo4 a bc def
a bc def

$ ./echo4 -s / a bc def
a/bc/def

$ ./echo4 -n a bc def
a bc def$

$ ./echo4 -help
Usage of ./echo4:
  -n    omit trailing newline
  -s string
		separator (default " ")

$ ./echo4 -s \* 1 2 3
1*2*3
*/

// 注意: 最后一个测试使用转移符(\)来保证分隔符能正常生效，否
// 则(*)不能作为分隔符。

// flag.Bool()函数创建了一个新的布尔标识变量。它有三个参数:
// (1)标识的名字("n")
// (2)变量的默认值(false)
// (3)当用户提供非法标识、非法参数抑或-h或--help参数时输出的消息
// flag.String()函数也使用名字、默认值、消息来创建一个字符串变量。
// 变量sep和n是指向标识变量的指针，它们必须通过*sep和*n来间接访问。

// 当程序运行时，在使用标识前，必须调用flag.Parse()来更新标识变量
// 的默认值。非标识参数也可以从flag.Args()返回的字符串切片来访问。
// 如果flag.Parse遇到错误，它输出一条帮助消息，然后调用
// os.Exit(2)来结束程序。
