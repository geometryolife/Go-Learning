// basename 函数模仿 UNIX shell 中的同名实用程序。只要 s 的前缀看起来
// 像是文件系统路径(各部分由斜杠分隔)，basename 就将其移除，文件类型
// 的后缀也被移除。
package main

import "fmt"

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}

// 初版本的 basename 独自完成全部工作，并不依赖任何库:
// basename 移除路径部分和 . 后缀
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// 将最后一个'/'和之前的部分全都舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个'.'之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

/*
>>> Execution Result:
c
c.d
abc
*/
