package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组 => 模拟队列
	front   int    // 表示指向队列的首部
	rear    int    // 表示指向队列的尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	// 判断队列是否已满，如果队满，则打印队满，返回
	if this.rear == this.maxSize-1 { // 注意: rear 是队列尾部(含最后的元素)
		return errors.New("The queue is full!")
	}

	// 如果队列未满，则可以继续添加元素，rear 后移，入队
	this.rear++
	this.array[this.rear] = val
	return
}

// 显示队列，找到队首，然后从队首遍历到队尾
func (this Queue) ShowQueue() {
	fmt.Println("队列当前的情况是: ")
	// this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 判断队列是否为空，如果 rear == front 则为空，返回
	if this.rear == this.front {
		return -1, errors.New("The queue is empty!")
	}

	// 队列未空，front 后移，元素出队
	this.front++
	val = this.array[this.front]
	return val, err
}

// 编写一个主函数测试
func main() {
	var key string
	var val int

	// 创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功!")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数:", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

// 总结:
// 1. 上述代码实现了基本的队列结构，但是没有有效地利用数组空间。
// 2. 请思考: 如何使用数组实现一个环形的队列?

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
加入队列成功!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
show
队列当前的情况是:
array[0]=1      array[1]=2      array[2]=3      array[3]=4      array[4]=5
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
6
The queue is full!
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
队列当前的情况是:

==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
add
输入要入队列的数:
8
The queue is full!
==>输入 add 添加数据到队列
==>输入 get 从队列获取元素
==>输入 show 显示队列
==>输入 exit 退出队列
exit
*/
