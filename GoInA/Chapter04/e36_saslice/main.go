// 设置长度和容量一样的好处
package main

import "fmt"

func main() {
	// 创建字符串切片，其长度和容量都是5个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 对第三个元素做切片,并限制容量，其长度和容量都是1个元素
	slice := source[2:3:3]
	fmt.Printf("%s\n", slice)
	fmt.Printf("%v\n", &slice[0])

	// 向slice追加新字符串
	slice = append(slice, "Kiwi")
	fmt.Printf("%s\n", slice)
	fmt.Printf("%v\n", &slice[0])
	// 通过限制容量，向新切片追加新元素时，不会修改原底层数组
	fmt.Printf("%s\n", source)

	// 如果不限制新切片的容量，那么向新切片追加新元素时，会修改底层数组
	fmt.Println()
	slice2 := source[2:3]
	fmt.Printf("%s\n", slice2)
	fmt.Printf("%v\n", &slice2[0])
	slice2 = append(slice2, "Kiwi")
	fmt.Printf("%s\n", slice2)
	fmt.Printf("%v\n", &slice2[0])
	fmt.Printf("%s\n", source)
}

/*
如果在创建切片时，设置切片的容量和长度一样，就可以强制让新切片的第一个
append 操作创建新的底层数组，与原有的底层数组分离。
>>> Execution Result:
[Plum]
0xc000078020
[Plum Kiwi]
0xc00000c080
[Apple Orange Plum Banana Grape]

[Plum]
0xc000078020
[Plum Kiwi]
0xc000078020
[Apple Orange Plum Kiwi Grape]
*/
