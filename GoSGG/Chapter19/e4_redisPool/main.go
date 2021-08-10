// Golang 操作 Redis 连接池
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的连接池
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		// 连接池中空闲连接的最大数目
		MaxIdle: 8,
		// 在给定时间由池分配的最大连接数，当为零时，池中的连接数没有限制
		MaxActive: 0,
		// 设置关闭连接的最大空闲时间，如果为零，则不关闭
		IdleTimeout: 100,
		// Dial 函数用于创建和配置连接
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

}
func main() {
	// 从池里取出一个连接
	conn := pool.Get()
	// 释放池使用的资源
	defer conn.Close()

	// Write
	_, err := conn.Do("SET", "name", "Joe")
	if err != nil {
		fmt.Println("SET err:", err)
		return
	}

	// Read
	r, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("GET err:", err)
	} else {
		fmt.Println(r)
	}

	// 如果要从连接池取出连接，一定要保证连接池没有被关闭
	// pool.Close()
	conn2 := pool.Get()
	fmt.Println("conn2", conn2)
	// Write
	_, err = conn2.Do("SET", "name2", "Tom")
	if err != nil {
		fmt.Println("SET err:", err)
		return
	}

	// Read
	r2, err := redis.String(conn2.Do("GET", "name2"))
	if err != nil {
		fmt.Println("GET err:", err)
	} else {
		fmt.Println(r2)
	}
}

/*
>>> Execution Result:
Joe
conn2 &{0xc0000982a0 0xc0000881e0 0}
Tom
*/
