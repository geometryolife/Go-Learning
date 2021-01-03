// 声明一个存储字符串切片的映射
package main

import "fmt"

func main() {
	// 创建一个映射，使用字符串切片作为值
	dict := map[int][]string{}

	for k, v := range dict {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
	fmt.Printf("映射dict的地址:     %p\n", dict)
	fmt.Printf("映射dict地址的地址: %p\n", &dict)
}

/*
映射的值为空，那么迭代不会返回键值对
>>> Execution Result:
映射dict的地址:     0xc000060150
映射dict地址的地址: 0xc00000e030
*/
