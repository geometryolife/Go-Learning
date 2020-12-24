// 同样类型的多维数组赋值
package main

import "fmt"

func main() {
	// 声明两个不同的二维整型数组
	var array1 [2][2]int
	var array2 [2][2]int

	// 为每个元素赋值
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40

	// 将array2的值复制给array1
	array1 = array2

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d,", array1[i][j])
		}
		fmt.Println()
	}
}

/*
>>> Execution Result:
10,20,
30,40,
*/
