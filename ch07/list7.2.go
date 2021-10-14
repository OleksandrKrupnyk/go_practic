// обслуживание файлов при помощи пользовательского обработчика
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", readme)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func readme(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ch07/files/readme.txt")
}
