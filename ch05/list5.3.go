// Листинг 5.3 Клиент сетевого регистратора
// // ncat.exe -lk 1902
package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902") //← Подключение к серверу  журналирования
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()              // ← Гарантировать закрытие соединения даже в случае аварии
	f := log.Ldate | log.Lshortfile // Оправка регистрационных сообщений ← в сетевое соединение
	logger := log.New(conn, "example ", f)
	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.") //← Вывести сообщение и инициировать
} // аварию, но не использовать Fatalln
