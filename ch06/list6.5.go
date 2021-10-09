// Буферизация вывода шаблона
package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./ch06/simple.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "My second page list 6.5",
		Content: "<h2>Content</h2> simple html code",
	}
	// Создание буфера для сохранения результатов обработки шаблона
	var b bytes.Buffer
	// Результат замены будет помещен в буфер
	err := t.Execute(&b, p)
	// При возникновении ошибки обработки шаблона будет выведена ошибка
	if err != nil {
		log.Fatal(" A error occured")
		return
	}
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
