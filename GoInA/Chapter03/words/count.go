// words 包提供了对单词计数的支持
package words

import "strings"

// CountWords 计算指定字符串中的单词数并返回计数
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
