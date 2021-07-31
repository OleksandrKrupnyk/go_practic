/**
List 3.10 Закрытие отправителем
*/
package main

import "time"

func main() {
	ch := make(chan bool)
	timeOut := time.After(2000 * time.Millisecond)
	go send(ch) //  Цикл вокруг select, проверяющего два канала и имеющего ветвь default
	for {
		select { // Получив сообщение по основному каналу,
		case <-ch: //  выведем что-нибудь определенное
			println("Got message")
		case <-timeOut: // По истечении заданного интервала
			println("Time out") //  времени завершим программу
			return
		default: // По умолчанию делаем небольшую паузу.
			println("*yawn*") // Это облегчает работу с примером
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	time.Sleep(500 * time.Millisecond)
	ch <- true
	close(ch)
	println("Send and closed")
}
