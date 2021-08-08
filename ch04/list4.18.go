// Листинг 4.18  Авария в функции response
package main

import (
	"bufio"
	"errors"
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

func response(data []byte, conn net.Conn) {
	panic(errors.New("Failure in response!")) //← Вместо выполнения каких-либо
} // полезных действий возбудим аварию
