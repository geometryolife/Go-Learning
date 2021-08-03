// 这个示例程序展示如何使用最基本的 log 包
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRECE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日记记录器
	log.Println("message")

	// Fatalln 在调用 Println() 之后会接着调用 os.Exit(1)
	log.Println("fatal message")

	// Panicln 在调用 Println() 之后会接着调用 panic()
	log.Panicln("panic message")
}

/*
>>> Execution Result:
TRECE: 2021/08/03 23:29:37.565433 /home/ubuntu/go/src/Go_Learning/GoInA/Chapter08/e1_listing03/main.go:15: message
TRECE: 2021/08/03 23:29:37.565503 /home/ubuntu/go/src/Go_Learning/GoInA/Chapter08/e1_listing03/main.go:18: fatal message
TRECE: 2021/08/03 23:29:37.565509 /home/ubuntu/go/src/Go_Learning/GoInA/Chapter08/e1_listing03/main.go:21: panic message
panic: panic message


goroutine 1 [running]:
log.Panicln(0xc000043f40, 0x1, 0x1)
        /usr/lib/go-1.13/src/log/log.go:352 +0xac
main.main()
        /home/ubuntu/go/src/Go_Learning/GoInA/Chapter08/e1_listing03/main.go:21 +0xe1
exit status 2
*/
