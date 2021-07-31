/**
List 3.9 Wrong close channel
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)
	go send(msg) // Вызов сопрограммы send с передачей канала для отправки
	for {
		select {
		case m := <-msg: //Если получено сообщение от send, вывести его
			fmt.Println(m)
		case <-until: // Выход по истечении заданного интервала времени.
			// Пауза позволит вам увидеть	ошибку, возникающую перед
			// выходом из сопрограммы main
			close(msg)
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan string) {
	for {
		ch <- "Hello!"
		time.Sleep(500 * time.Millisecond)
	}
}
