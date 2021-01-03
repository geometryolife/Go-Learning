// 使用make声明映射
package main

import "fmt"

func main() {
	// 创建一个映射，键的类型是string，值的类型是int
	dict1 := make(map[string]int)

	// 创建一个映射，键和值的类型都是string
	// 使用两个键值对初始化映射
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	for k, v := range dict1 {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	for k, v := range dict2 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
}

/*
dict1声明了一个空映射，所以迭代没有返回任何键值对
>>> Execution Result:
Key: Red, Value: #da1337
Key: Orange, Value: #e95a22
*/
