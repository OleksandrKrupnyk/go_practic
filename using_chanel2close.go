/**
List 3.11 Использование завершающего канала
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)              //  Дополнительный канал с типом данных bool
	until := time.After(5 * time.Second) //  для сообщения о завершении
	go send2(msg, done)                  // Передача двух каналов в send

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until: //  По истечении заданного интервала времени сообщить
			done <- true // сопрограмме send, что работа завершена
			time.Sleep(500 * time.Millisecond)
			return

		}
	}
}

func send2(ch chan<- string, done <-chan bool) { // ch используется для отправки,
	for { // а done – для получения
		select {
		case v := <-done: // Завершить работу после получения сообщения
			print("Done", v) // из канала done
			close(ch)
			return
		default:
			ch <- "Hello"
			time.Sleep(500 * time.Millisecond)
		}

	}
}
