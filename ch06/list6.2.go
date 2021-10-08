// Использование простого ХМЛ шаблона
package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "My fist Page",
		Content: "I lose my favorite game <b>Bob Dilan</b>",
	}
	t := template.Must(template.ParseFiles("./ch06/simple.html"))
	err := t.Execute(w, p)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", displayPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
