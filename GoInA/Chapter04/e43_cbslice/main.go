// 组合切片的切片
package main

import "fmt"

func main() {
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100, 200}}

	fmt.Printf("slice:        %p\n", slice)
	fmt.Printf("&slice:       %p\n", &slice)
	fmt.Printf("slice[0]:     %p\n", slice[0])
	fmt.Printf("slice[1]:     %p\n", slice[1])
	fmt.Printf("&slice[0]:    %p\n", &slice[0])
	fmt.Printf("&slice[1]:    %p\n", &slice[1])
	fmt.Printf("&slice[0][0]: %p\n", &slice[0][0])
	fmt.Printf("&slice[1][0]: %p\n", &slice[1][0])
	fmt.Printf("&slice[1][1]: %p\n", &slice[1][1])

	// 为第一个切片追加值为20的元素
	slice[0] = append(slice[0], 20)

	fmt.Println()
	fmt.Println("切片的地址")
	fmt.Printf("slice:        %p\n", slice)
	fmt.Println("切片的地址的地址")
	fmt.Printf("&slice:       %p\n", &slice)
	fmt.Println("第一个内层切片底层数组的首地址")
	fmt.Printf("slice[0]:     %p\n", slice[0])
	fmt.Println("第二个内层切片底层数组的首地址")
	fmt.Printf("slice[1]:     %p\n", slice[1])
	fmt.Println("第一个内层切片的地址")
	fmt.Printf("&slice[0]:    %p\n", &slice[0])
	fmt.Println("第二个内层切片的地址")
	fmt.Printf("&slice[1]:    %p\n", &slice[1])
	fmt.Println("元素下标为[0][0]的底层数组的地址")
	fmt.Printf("&slice[0][0]: %p\n", &slice[0][0])
	fmt.Println("元素下标为[0][1]的底层数组的地址")
	fmt.Printf("&slice[0][1]: %p\n", &slice[0][1])
	fmt.Println("元素下标为[1][0]的底层数组的地址")
	fmt.Printf("&slice[1][0]: %p\n", &slice[1][0])
	fmt.Println("元素下标为[1][1]的底层数组的地址")
	fmt.Printf("&slice[1][1]: %p\n", &slice[1][1])
}
