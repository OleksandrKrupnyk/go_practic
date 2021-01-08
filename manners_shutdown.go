/*
List 2.15
*/
package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

/*
Тип даннх
*/
type handler struct{}

/*
Функция возвращает ссылку на пустую структуру
*/
func newHandler() *handler {
	return &handler{}
}

/*
Функця, реализует интерфейс
*/
func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Oleksandr Krupnyk"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

/*
Горутина
*/
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

func main() {

	handler := newHandler()
	// Канал котрый передает сигналы ОС
	ch := make(chan os.Signal)

	// Настройка мониторинга сигналов операционной системы
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	// Запус веб сервера
	manners.ListenAndServe("127.0.0.1:8080", handler)
}
