package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	queueList()
}

func queueList() {
	// Connection
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err:", err)
		return
	} else {
		fmt.Println("Connected successfully!")
	}

	defer conn.Close()

	// Write
	_, err = conn.Do("RPUSH", "books", "golang", "java", "python")
	if err != nil {
		fmt.Println("RPUSH err:", err)
		return
	} else {
		fmt.Println("Write successfully!")
	}

	// Read
	r1, err := conn.Do("LLEN", "books")
	if err != nil {
		fmt.Println("LLEN err:", err)
		return
	} else {
		fmt.Println("List Length:", r1)
	}

	// Dequeue
	for {
		r2, err := redis.String(conn.Do("LPOP", "books"))
		if err != nil {
			fmt.Println("LPOP err:", err)
			return
		}
		fmt.Println(r2)
	}
}

/*
>>> Execution Result:
Connected successfully!
Write successfully!
List Length: 3
golang
java
python
LPOP err: redigo: nil returned
*/
