// Echo3 prints its command-lind arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0:])
	fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	// fmt.Println(os.Args[2:])
}
