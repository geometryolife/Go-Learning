package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	singleHash()
	singleHash2()
	multiHash()
}

func singleHash() {
	// Connection
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected successfully!")

	// Write
	_, err = conn.Do("HSET", "user1", "name", "Joe Chen")
	if err != nil {
		fmt.Println("HSET err =", err)
		return
	}

	_, err = conn.Do("HSET", "user1", "age", 30)
	if err != nil {
		fmt.Println("HSET err =", err)
		return
	}
	fmt.Println("Write successfully!")

	// Read
	r1, err := conn.Do("HGET", "user1", "name")
	if err != nil {
		fmt.Println("HGET err =", err)
		return
	}

	r2, err := conn.Do("HGET", "user1", "age")
	if err != nil {
		fmt.Println("HGET err =", err)
		return
	}

	// 如果不使用类型转换，则无法显示正确的类型
	fmt.Printf("Read successfully!\n%T\t%s\n%T\t%d\n\n", r1, r1, r2, r2)
}

func singleHash2() {
	// Connection
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected successfully!")

	// Write
	_, err = conn.Do("HSET", "user2", "name", "Joe Chen")
	if err != nil {
		fmt.Println("HSET err =", err)
		return
	}

	_, err = conn.Do("HSET", "user2", "age", 20)
	if err != nil {
		fmt.Println("HSET err =", err)
		return
	}
	fmt.Println("Write successfully!")

	// Read
	r1, err := redis.String(conn.Do("HGET", "user2", "name"))
	if err != nil {
		fmt.Println("HGET err =", err)
		return
	}

	r2, err := redis.Int(conn.Do("HGET", "user2", "age"))
	if err != nil {
		fmt.Println("HGET err =", err)
		return
	}
	fmt.Printf("Read successfully!\n%T\t%v\n%T\t%v\n\n", r1, r1, r2, r2)
}

func multiHash() {
	// Connection
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected successfully!")

	// Write
	_, err = conn.Do("HMSET", "user3", "name", "贝尔·格里尔斯", "age", 47,
		"country", "UK")
	if err != nil {
		fmt.Println("HMSET err =", err)
		return
	}
	fmt.Println("Write successfully!")

	// Read
	r1, err := redis.Strings(conn.Do("HMGET", "user3", "name", "age"))
	if err != nil {
		fmt.Println("HMGET err =", err)
		return
	}
	fmt.Printf("Read successfully!\n%T\t%v\n", r1, r1)

	for i, v := range r1 {
		fmt.Printf("%T\tr1[%d] = %s\n", v, i, v)
	}

	fmt.Println()

	r2, err := redis.Values(conn.Do("HMGET", "user3", "name", "age"))
	if err != nil {
		fmt.Println("HMGET err =", err)
		return
	}
	fmt.Printf("Read successfully!\n%T\t%v\n", r2, r2)

	for i, v := range r2 {
		fmt.Printf("%T\tr2[%d] = %s\n", v, i, v)
	}

	fmt.Println()

	// Read All
	r3, err := redis.Strings(conn.Do("HGETALL", "user3"))
	if err != nil {
		fmt.Println("HGETALL err =", err)
		return
	} else {
		fmt.Printf("Read successfully!\n%T\t%v\n", r3, r3)
	}

	for i, v := range r3 {
		fmt.Printf("%T\tr3[%d] = %s\n", v, i, v)
	}
	for i := 0; i < len(r3); i++ {
		if i%2 != 0 {
			fmt.Printf("%s: %s\n", r3[i-1], r3[i])
		}
	}
}

/*
>>> Execution Result:
Connected successfully!
Write successfully!
Read successfully!
[]uint8 Joe Chen
[]uint8 [51 48]

Connected successfully!
Write successfully!
Read successfully!
string  Joe Chen
int     20

Connected successfully!
Write successfully!
Read successfully!
[]string        [贝尔·格里尔斯 47]
string  r1[0] = 贝尔·格里尔斯
string  r1[1] = 47

Read successfully!
[]interface {}  [[232 180 157 229 176 148 194 183 230 160 188 233 135 140 229 176 148 230 150 175] [52 55]]
[]uint8 r2[0] = 贝尔·格里尔斯
[]uint8 r2[1] = 47

Read successfully!
[]string        [name 贝尔·格里尔斯 age 47 country UK]
string  r3[0] = name
string  r3[1] = 贝尔·格里尔斯
string  r3[2] = age
string  r3[3] = 47
string  r3[4] = country
string  r3[5] = UK
name: 贝尔·格里尔斯
age: 47
country: UK
*/
