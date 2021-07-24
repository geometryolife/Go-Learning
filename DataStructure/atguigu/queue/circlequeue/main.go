package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int // 数组 => 模拟队列
	head    int    // 指向队列的队首
	tail    int    // 指向队列队队尾
}

// 入队方法 Push()
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("The queue is full!")
	}
	// 入队后，this.tail 在队列尾部，但所指向的位置不包含元素
	this.array[this.tail] = val // 把值传给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队方法 Pop()
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("The queue is empty!")
	}
	// 取出 this.head 的元素，head 指向对首，并且所指位置包含元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// 显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下: ")
	// 计算当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("The queue is empty!")
	}
	// 增加一个辅助变量，值向 head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断环形队列为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 计算环形队列有多少个元素
func (this *CircleQueue) Size() int {
	// 这是一个关键的算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	var key string
	var val int

	// 初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	for {
		fmt.Println("==>输入 add 添加数据到队列")
		fmt.Println("==>输入 get 从队列获取元素")
		fmt.Println("==>输入 show 显示队列")
		fmt.Println("==>输入 exit 退出队列")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Println("输入要入队列的数: ")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功!")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数:", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

/*
>>> Execution Result:
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
1
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
2
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
3
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
4
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
5
The queue is full!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
show
环形队列情况如下:
arr[0]=1        arr[1]=2        arr[2]=3        arr[3]=4
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
get
从队列中取出了一个数: 1
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
show
环形队列情况如下:
arr[1]=2        arr[2]=3        arr[3]=4
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
5
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
show
环形队列情况如下:
arr[1]=2        arr[2]=3        arr[3]=4        arr[4]=5
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
get
从队列中取出了一个数: 2
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
get
从队列中取出了一个数: 3
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
get
从队列中取出了一个数: 4
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
get
从队列中取出了一个数: 5
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
get
The queue is empty!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
show
环形队列情况如下:
The queue is empty!

==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
6
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
show
环形队列情况如下:
arr[0]=6
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
exit
*/
