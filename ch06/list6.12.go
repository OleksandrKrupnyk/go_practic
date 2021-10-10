// Использование наследования шаблонов
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Объявление переменная для хранения списка шаблонов
var t2 map[string]*template.Template

func init() {
	// Инициализация памяти для переменной
	t2 = make(map[string]*template.Template)
	// Загрузка шаблонов вместе с основным шаблоном
	temp := template.Must(template.ParseFiles("./ch06/base.gohtml", "./ch06/user.gohtml"))
	// сохранить ссылку в карту
	t2["user.gohtml"] = temp
	temp = template.Must(template.ParseFiles("./ch06/base.gohtml", "./ch06/page.gohtml"))
	t2["page.gohtml"] = temp
}

// PageT Структура для хранения
type PageT struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPageT(w http.ResponseWriter, r *http.Request) {
	log.Printf("User go to page @base@ %s", r.Header)
	p := &PageT{
		Title:   "Base page title",
		Content: "Some interesting page content",
	}
	err := t2["page.gohtml"].ExecuteTemplate(w, "base", p)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("User go to page @user@ %s", r.Header["Referer"])
	u := &User{
		Name:     "Sasha",
		Username: "Unrealer2004",
	}
	err := t2["user.gohtml"].ExecuteTemplate(w, "base", u)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}
// Обработка запроса получения иконки
func displayFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"./ch06/favicon.ico")
	return
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPageT)
	http.HandleFunc("/favicon.ico", displayFavicon)
	http.ListenAndServe(":8080", nil)
}
