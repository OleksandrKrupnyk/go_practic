//Листинг 4.19 Обработка аварий в сопрограмме
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
	// Отложенная функция обрабатывает
	//аварию и гарантирует закрытие
	//соединения в любом случае
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s \n", err)
		}
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket.")
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	fmt.Println(string(data))
	_, err := conn.Write(data)
	if err != nil {
		return
	}
	panic(errors.New("pretend I'm a real error")) //← И снова возбудить аварию для имитации сбоя
}
