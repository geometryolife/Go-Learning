// 声明二维数组
package main

import "fmt"

func main() {
	// 声明一个二维整型数组，两个维度分别存储4个元素和2个元素
	var array1 [4][2]int

	// 使用数组字面量来声明并初始化一个二维整型数组
	array2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	// 声明并初始化外层数组中索引为1和3的元素
	array3 := [4][2]int{1: {20, 21}, 3: {40, 41}}

	// 声明并初始化外层数组和内层数组的单个元素
	array4 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Printf("array1:")
	for i := 0; i < 4; i++ {
		fmt.Println()
		for j := 0; j < 2; j++ {
			fmt.Printf("%d, ", array1[i][j])
		}
	}

	fmt.Printf("\narray2:")
	for i := 0; i < 4; i++ {
		fmt.Println()
		for j := 0; j < 2; j++ {
			fmt.Printf("%d, ", array2[i][j])
		}
	}

	fmt.Printf("\narray3:")
	for i := 0; i < 4; i++ {
		fmt.Println()
		for j := 0; j < 2; j++ {
			fmt.Printf("%d, ", array3[i][j])
		}
	}

	fmt.Printf("\narray4:")
	for i := 0; i < 4; i++ {
		fmt.Println()
		for j := 0; j < 2; j++ {
			fmt.Printf("%d, ", array4[i][j])
		}
	}
}

/*
>>> Execution Result:
array1:
0, 0,
0, 0,
0, 0,
0, 0,
array2:
10, 11,
20, 21,
30, 31,
40, 41,
array3:
0, 0,
20, 21,
0, 0,
40, 41,
array4:
0, 0,
20, 0,
0, 0,
0, 41,
*/
