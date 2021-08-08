// Листинг 4.20    Небольшой HTTP-сервер
package main

import (
	"errors"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) //← Передача функции handler HTTP-серверу
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	} // ← Запуск сервера
}
func handler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	panic(errors.New("fake panic"))
}
