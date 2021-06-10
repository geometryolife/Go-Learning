package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	// 源码中，如果在浮点数或十进制整数后面紧接着写字母 i，如 2i 或
	// 3.141592i，它就变成一个虚数，表示一个实部为0的复数:
	fmt.Println(1i * 1i) // "(-1+0i)", i² = -1

	a := 1 + 2i
	b := 3 + 4i
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)

	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"
}

/*
>>> Execution Result:
(1+2i)
(3+4i)
(-5+10i)
-5
10
(-1+0i)
(1+2i)
(3+4i)
(4+6i)
(0+1i)
*/
