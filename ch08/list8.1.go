// Приклад використення функції http.Get
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, _ := http.Get("http://bispinor.net")
	b, _ := ioutil.ReadAll(res.Body)
	err := res.Body.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	file, err := os.OpenFile("./index.html", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}(file)
	_, err = file.Write(b)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
