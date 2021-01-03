// 从映射获取值并判断键是否存在
package main

import "fmt"

func main() {
	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}

	// 将Blue的代码加入到映射
	colors["Blue"] = "#0000FF"

	// 获取键Blue对应的值
	value, exists := colors["Blue"]
	// 获取键Coral对应的值
	v, e := colors["Coral"]

	// 这个键存在吗?
	// 存在
	if exists {
		fmt.Println(value)
		fmt.Println(exists)
	}
	// 不存在
	// 由于映射里未加入Coral的颜色代码，键Coral不存在映射里，e的值为false，
	// v的值为空字符串
	if !e {
		fmt.Println(v)
		fmt.Println(e)
	}

}

/*
通过键来索引映射时，即便这个键不存在也总会返回一个值。
在这种情况下返回的是该值对应的类型的零值。
>>> Execution Result:
#0000FF
true

false
*/
