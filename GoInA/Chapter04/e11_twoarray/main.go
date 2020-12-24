// 访问二维数组的元素
package main

import "fmt"

func main() {
	// 声明一个2*2的二维数组
	var array [2][2]int

	// 设置每个元素的整型值
	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d,", array[i][j])
		}
		fmt.Println()
	}
}

/*
>>> Execution Result:
10,20,
30,40,
*/
