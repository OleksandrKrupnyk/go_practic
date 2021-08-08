// Листинг 5.1 Простейшее использование пакета log
package main

import (
	"log"
)

func main() {
	log.Println("This is a regular message.")       //← Запись сообщения в os.Stderr
	log.Fatalln("This is a fatal error.")           // Запись сообщения в os.Stderr и выход с кодом ошибки
	log.Println("This is the end of the function.") // Эта инструкция никогда не будет выполнена
}
