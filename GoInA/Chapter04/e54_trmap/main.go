// 在函数间传递映射
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

	// 调用函数来移除指定的键
	removeColor(colors, "Coral")
	fmt.Println()

	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

}

// removeColor将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

/*
>>> Execution Result:
Key: AliceBlue  Value: #F0F8FF
Key: Coral  Value: #FF7f50
Key: DarkGray  Value: #A9A9A9
Key: ForestGreen  Value: #228B22

Key: AliceBlue  Value: #F0F8FF
Key: DarkGray  Value: #A9A9A9
Key: ForestGreen  Value: #228B22
*/
