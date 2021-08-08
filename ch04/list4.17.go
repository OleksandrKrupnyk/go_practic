// Листинг 4.17 Эхо-сервер
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026") // ← Запуск нового сервера, обслуживающего порт 1026
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handle(conn) //← При появлении запроса передать его в функцию handle
	}
}
func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket.")
		conn.Close()
	}
	response(data, conn) //← При получении строки передать ее в функцию response
}

// Прием новых клиентских запросов и обработка ошибок подключения
//Попытка чтения строки из подключения
//В случае ошибки чтения строки вывести сообщение и закрыть подключение
func response(data []byte, conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	_, err := conn.Write(data)
	if err != nil {
		return
	}
}
