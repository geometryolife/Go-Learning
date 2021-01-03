// 使用range迭代映射
package main

import "fmt"

func main() {
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#F0F8FF",
		"Coral":       "#FF7f50",
		"DarkGray":    "#A9A9A9",
		"ForestGreen": "#228B22",
	}

	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
}

/*
>>> Execution Result:
Key: AliceBlue  Value: #F0F8FF
Key: Coral  Value: #FF7f50
Key: DarkGray  Value: #A9A9A9
Key: ForestGreen  Value: #228B22
*/
