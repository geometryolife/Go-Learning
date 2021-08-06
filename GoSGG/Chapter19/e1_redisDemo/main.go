// Set/Get 接口
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 通过 Go 向 Redis 写入数据和读取数据
	// 1. 连接到 Redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	defer conn.Close()
	fmt.Println("connection succeeded!", conn)

	// 2. 通过 Go 向 Redis 写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "贝尔·格里尔斯")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	fmt.Println("Successfully written!")

	// 3. 通过 Go 向 Redis 读取数据 string [key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	// redis 包内部提供了转换函数，可以使用自带的转换函数。
	// 转换错误示例:
	// 因为返回的 r 是 interface{}，而 name 对应的值是 string，因此需要转换
	// nameString := r.(string)

	fmt.Printf("Read successfully!\n%s\n", r)
}

/*
>>> Execution Result:
connection succeeded! &{{0 0} 0 <nil> 0xc000010038 0 0xc000056240 0 0xc000054
080 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0
 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
Successfully written!
Read successfully!
贝尔·格里尔斯
*/
