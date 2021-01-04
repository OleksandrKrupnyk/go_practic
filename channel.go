/*
List 1.7
*/
package main

import (
	"fmt"
	"time"
)

/*
Считывать с канала число и читать
*/
func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	// Создание канала
	c := make(chan int)
	// Массив
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	// запуск горутины
	go printCount(c)
	// передача значения в канал
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of Main")
}
