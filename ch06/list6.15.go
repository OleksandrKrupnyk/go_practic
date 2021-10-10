// объединение шаблонов
package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

// Переменная для хранения шаблона
var t8 *template.Template

// Переменная для хранения безопасной строки
var qc template.HTML

func init() {
	t8 = template.Must(template.ParseFiles("./ch06/index.gohtml", "./ch06/quote.gohtml"))
}

// PageD Структура для хранения содержимого страницы
type PageD struct {
	Title   string
	Content template.HTML
}

// Quote структура для хранения цитаты
type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote:  "Не лезь вводу не зная броду",
		Person: "Народная мудрость",
	}
	var b bytes.Buffer
	err := t8.ExecuteTemplate(&b, "quote.gohtml", q)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	qc = template.HTML(b.String())

	http.HandleFunc("/", displayQuote)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

//displayQuote Обработчик страницы
func displayQuote(w http.ResponseWriter, r *http.Request) {
	p := &PageD{
		Title:   "Страничка народных мудростей",
		Content: qc,
	}
	err := t8.ExecuteTemplate(w, "index.gohtml", p)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
