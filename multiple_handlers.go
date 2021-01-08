/*
List 2.16
*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello2)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage3)
	http.ListenAndServe(":8080", nil)

}

/*
Страница прощания
*/
func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Oleksandr Krupnyk"
	}
	fmt.Fprint(res, "Goodbye my friend ", name)
}

/*
Домашняя страница
*/
func homePage3(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "This is home page")
}

/*
Страница приветствия
*/
func hello2(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Oleksandr Krupnyk"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}
