// использование пакета path для определения путей
package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

/*
хеш таблица обработчиков
*/
type pathResolver struct {
	// тип хранимых данных
	handlers map[string]http.HandlerFunc
}

/*
Новый экземпляр структуры
*/
func newPathResolver() *pathResolver {
	// Заполнение структуры
	return &pathResolver{
		handlers: make(map[string]http.HandlerFunc),
	}
}

// Add Добавление обработчика в хеш таблицу
func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

// hello
func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Oleksandr Krupnyk"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

/*
Реализация интерфейса обслуживания сервера
*/
func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Объединение метода и пути для проверки
	check := req.Method + " " + req.URL.Path
	// Паттерн=Обработчик
	for pattern, handlerFunc := range p.handlers {
		// Выполнение действия; проверка
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			log.Fatalln(err.Error())
		}
	}
	// Если пути не нашлось
	http.NotFound(res, req)
}

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)

	dir := http.Dir("./ch07/files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	pr.Add("GET /static/*", handler.ServeHTTP)

	err := http.ListenAndServe(":8080", pr)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
