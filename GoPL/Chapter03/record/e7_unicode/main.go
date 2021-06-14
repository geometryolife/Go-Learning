package main

import "fmt"

func main() {
	s := "All is well with the world."
	a := "world"
	b := "All"
	c := "."

	fmt.Println(HasPrefix(s, a))
	fmt.Println(HasSuffix(s, a))
	fmt.Println(Contains(s, a))

	fmt.Println()
	fmt.Println(HasPrefix(s, b))
	fmt.Println(HasSuffix(s, c))
}

// 判断某个字符串是否包含另一个字符串的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 判断某个字符串是否包含另一个字符串的后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 判断某个字符串是否包含另一个子字符串
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

/*
>>> Execution Result:
false
false
true

true
true
*/
