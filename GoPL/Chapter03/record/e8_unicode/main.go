package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"

	// s 包含13个字节，按 UTF-8 解读，本质是9个码点或文字符号的编码
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	fmt.Println()

	// UTF-8 解码器
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println()

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println()

	// 用 range 循环统计字符串中的文字符号数目
	n1 := 0
	// 这条表达式是冗余的
	for _, _ = range s {
		n1++
	}

	// 简化 range 表达式，忽略没用的变量
	n2 := 0
	for range s {
		n2++
	}

	fmt.Println(n1, n2)
}

/*
>>> Execution Result:
13
9

0       H
1       e
2       l
3       l
4       o
5       ,
6
7       世
10      界

0       'H'     72
1       'e'     101
2       'l'     108
3       'l'     108
4       'o'     111
5       ','     44
6       ' '     32
7       '世'    19990
10      '界'    30028

9 9
*/
