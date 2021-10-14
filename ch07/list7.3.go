// обслуживание подкаталога
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title   string
	Content string
}

func main() {
	dir := http.Dir("/home/sasha")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handler)
	http.HandleFunc("/", displayPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
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
