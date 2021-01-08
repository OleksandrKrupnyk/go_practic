/*
List 2.18
*/
package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

/*
Тип данных для хранения хеш таблицы рег.выражение=обработчика
*/
type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

/*
Функция заполнения заполнения структуры
*/
func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

func main() {
	// Структура
	rr := newPathResolver()
	// Добавление данных в хеш=таблицу
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodbye)
	// Запуск сервера
	http.ListenAndServe(":8080", rr)
}

/*
Метод добавления данных в структуру
*/
func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	// рег.выражение = обработчик
	r.handlers[regex] = handler
	// определение  кеша
	cache, _ := regexp.Compile(regex)
	// сохранение кеша что бы каждый раз не вычислять
	r.cache[regex] = cache
}

/*
Реализация интерфейса а именно метода ServeHTTP
*/
func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Поиск и вызов обработчика
	check := req.Method + " " + req.URL.Path
	// Паттерн=Обработчик
	for pattern, handlerFunc := range r.handlers {
		// Сравнение кеша из хеш-таблицы
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}
	// Обработчик по умолчанию
	http.NotFound(res, req)
}

/*
Обработчик страницы приветствия
*/
func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Олександр Крупник"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

/*
Обработчик страницы прощания
*/
func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Олександр Крупник"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
