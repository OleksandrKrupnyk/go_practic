// использование вложених шаблонов
package main

import (
	"html/template"
	"net/http"
)
// Переменная для хранения объекта шаблона
var t *template.Template
func init() {
	// чтение двух файлов шаблонов
	t = template.Must(t.ParseFiles("./ch06/simple.gohtml", "./ch06/head.gohtml"))
}
// Структура для передачи переменых в шаблоны
type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Include HTML template",
		Content: "<b>Html</b> - is hyper text makeup language",
	}
	// Обработка шаблона с передачей данных
	// Указывается имя шаблона(файла) который является главным для всех шаблонов
	t.ExecuteTemplate(w,"simple.gohtml", p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
