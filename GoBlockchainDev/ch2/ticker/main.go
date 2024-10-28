package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("发射！")
}

func main() {
	ticker := time.NewTicker(time.Second)
	num := 5
	for {
		<- ticker.C
		fmt.Println(num)
		num--
		if num == 0 {
			break
		}
	}
	ticker.Stop()
	launch()
}
