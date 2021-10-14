// страницы ошибок для пользоваельского файл сервера
package main

import (
	"fmt"
	fs "github.com/Masterminds/go-fileserver"
	"log"
	"net/http"
)

func main() {
	fs.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "The request page could not be found.")
	}
	dir := http.Dir("./ch07/files")
	err := http.ListenAndServe(":8080", fs.FileServer(dir))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
