/*
List 2.17
*/
package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
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

/*
Добавление обработчика в хеш таблицу
*/
func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

/*
Обработчик приветствия
*/
func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Oleksandr Krupnyk"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

/*
Обработчик прощания
/goodbye/<name>
*/
func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Oleksandr Krupnyk"
	}
	fmt.Fprint(res, "Goodbye ", name)
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
			fmt.Fprint(res, err)
		}
	}
	// Если пути не нашлось
	http.NotFound(res, req)
}

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":8080", pr)
}
