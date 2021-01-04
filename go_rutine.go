/**
List 1.6
*/
package main

import (
	"fmt"
	"time"
)

/**
Counting form zero to five
*/
func cont() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go cont()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 2)
}
