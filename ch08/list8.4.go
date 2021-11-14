package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// hasTimedOut - функция возвращает true если ошибка вызвана превышение времени ожидания
func hasTimedOut(err error) bool {
	switch err := err.(type) {
	// Ошибка url.Error может возникать из-за превышения времени ожидания в пакете net
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
		// Проверить не обнаружил ли ошибку привышения ожидания пакет net
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}

	}
	// превышение времени ожидания может быть выявлено дополнительной проверкой
	// некоторых ошибок
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

func main() {
	_, err := http.Get("http://example.com/test.zip")
	if err != nil && hasTimedOut(err) {
		log.Println("A time error occurred")
	}
	if err != nil {
		log.Println(err.Error())
	}
}
