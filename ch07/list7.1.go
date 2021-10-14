// Обслуживание фалов при помощи пакета http
package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.Dir("/home/sasha")
	err := http.ListenAndServe(":8080", http.FileServer(dir))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
