// внедрение файлов в испольняемый файл при помощи go.rice
package main

import (
	"github.com/GeertJohan/go.rice"
	"log"
	"net/http"
)

func main() {
	box := rice.MustFindBox("../ch07/files/")
	httpbox := box.HTTPBox()
	err := http.ListenAndServe(":8080", http.FileServer(httpbox))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
