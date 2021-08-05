/**
List 3.12  Простая блокировка посредством каналов
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1) // Создание буферизированного канала
	for i := 1; i < 70; i++ {  // Вызов шести сопрограмм, совместно использующих блокирующий канал
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)

}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants to lock channel\n", id)
	lock <- true //Рабочий процесс устанавливает блокировку, отправляя сообщение.
	//Первый рабочий процесс захватывает
	//единичный объем, что делает его
	//собственником блокировки. Остальные
	//окажутся заблокированными
	fmt.Printf("%d has the lock channel\n", id)
	time.Sleep(500 * time.Millisecond) //  Фрагмент между lock <- true и <- lock выполняется под защитой блокировки
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock // Снять блокировку, прочитав значение из канала. В результате в буфере
	//  освободится место, и следующая функция сможет установить блокировку
}
