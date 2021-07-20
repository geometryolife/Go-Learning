package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1. 创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 蓝子

	// 2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	fmt.Println()

	// 3. 转成稀疏数组。想 -> 算法
	// 思路:
	// (1). 遍历 chessMap，如果发现一个元素不为0，则创建一个 valNode 结构体
	// (2). 将其放入到对应的切片里即可

	var sparseArr []ValNode

	// 标准的稀疏数组还含有一条记录表示原始二维数组的规模(行、列、默认值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个 valNode 值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d, %d, %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	fmt.Println()

	// 直接使用稀疏数组恢复原始的数组
	// 先创建一个原始数组
	var chessMap2 [11][11]int

	// 然后遍历 sparseArr
	for i, valNode := range sparseArr {
		if i != 0 { // 跳过第1行记录值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	// 打印查看恢复后的原始数组 chessMap2
	fmt.Println("恢复后的原始数组......")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
