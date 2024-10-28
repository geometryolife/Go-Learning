package main

import (
	"fmt"
	"time"
)

var c1 chan int
var c2 chan int

func main() {
	c1 = make(chan int)
	c2 = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			num := <- c1
			c2 <- num * num
		}
	}()

	for {
		num := <- c2
		fmt.Println(num)
	}
}
