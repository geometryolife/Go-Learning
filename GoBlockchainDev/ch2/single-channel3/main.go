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

	go func(out chan <- int) {
		for i := 0; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second * 1)
		}
		close(out)
	}(c1)

	go func(in <- chan int, out chan <- int) {
		for {
			num, ok := <- in
			if ok {
				out <- num * num
			} else {
				break
			}
		}
		close(out)
	}(c1, c2)

	for {
		num, ok := <- c2
		if ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}
