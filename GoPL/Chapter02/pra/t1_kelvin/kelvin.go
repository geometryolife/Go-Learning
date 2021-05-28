// 练习2.1: 添加类型、常量和函数到 tempconv 包中，处理以开尔文为
// 单位(K)的温度值，0K=-273.15°C，变化 1K 和变化 1°C 是等价的。
package main

import (
	"Go_Learning/GoPL/Chapter02/demo/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		k := tempconv.Kelvin(t)
		// c := tempconv.Celsius(t)
		fmt.Printf("%s = %s\n", k, tempconv.KToC(k))
	}
}
