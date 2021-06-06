package main

import (
	"fmt"
	"math"
)

func main() {
	booboo()
	fmt.Println()
	natural()
	fmt.Println()
	special()
}

func booboo() {
	// float32 能精确表示的正整数范围有限
	var f float32 = 16777216 // 1 << 24
	var f2 float32 = 1 << 24
	fmt.Printf("%b\n%b\n", f, f2)
	fmt.Println(f == f+1)

	// 在源码中，浮点数可以写成小数
	const e = 2.71828 // (e 的近似值)

	// 科学记数法
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	fmt.Println(e)
	fmt.Println(Avogadro)
	fmt.Println(Planck)
}

func natural() {
	for n := 0; n < 8; n++ {
		fmt.Printf("n = %d    eⁿ = %8.3f\n", n, math.Exp(float64(n)))
	}
}

func special() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
