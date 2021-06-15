// 1. 在包级别，声明的顺序和它们的作用域没有关系，所以一个声明可以
// 引用它自己或者跟在它后面的其他声明，使我们可以声明递归或相互递归
// 的类型和函数。如果常量或变量声明引用它自己，则编译器会报错。
// 2. 在 if 块中处理错误然后返回，这样成功执行的路径不会被变得支离破碎。
// 3. 短变量声明依赖一个明确的作用域。
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		if f, err := os.open(fname); err != nil { // 编译错误: 未使用 f
			return err
		}
		f.stat()  // 编译错误: 未定义 f
		f.close() // 编译错误: 未定义 f
	*/

	// 通常需要在条件判断之前声明 f，使其在 if 语句后面可以访问
	/*
		f, err := os.Open(fname)
		if err != nil {
			return err
		}
		f.Stat()
		f.Close()
	*/

	// 避免在外部块中声明 f 和 err，方法是将 Stat 和 Close 的
	// 调用放到 else 块中:
	/*
		if f, err := os.Open(fname); err != nil {
			return err
		} else {
			// f 与 err 在这里可见
			f.Stat()
			f.Close()
		}
	*/

	fmt.Println(cwd)
}

var cwd string

// 获取当前的工作目录，然后把它保存在一个包级别的变量里
// 因为 cwd 和 err 在 init 函数块的内部都尚未声明，所以 := 语句将
// 它们都声明为局部变量。内层 cwd 的声明让外部的声明不可见，所以
// 这个语句没有按预期更新包级别的 cwd 变量:
/*
func init() {
	cwd, err := os.Getwd() // 编译错误: 未使用 cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
*/

// 此时，全局的 cwd 变量依然未初始化，看起来一个普通的日志输出
// 让 bug 变得不明显:

func init() {
	cwd, err := os.Getwd() // 注意: 错误
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	// 增加引用局部 cwd 变量的日志语句就可以让检查失效
	log.Printf("Working directory = %s", cwd)
}

// 处理上述问题的方法有许多，最直接的方法是在另一个 var 声明中
// 声明 err，避免使用 :=
func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

/*
>>> Execution Result:
2021/06/15 20:35:02 Working directory = /home/ubuntu/go/src/Go_Learning/GoPL/Chapter02/record/e19_scope
/home/ubuntu/go/src/Go_Learning/GoPL/Chapter02/record/e19_scope
*/
