// Простий користувацький кліент
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	cc := http.Client{Timeout: time.Millisecond * 500}
	res, err := cc.Get("http://bispinor.net")
	if err != nil {
		log.Fatalln(err.Error())
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = res.Body.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	file, err := os.OpenFile("./index2.html", os.O_CREATE|os.O_WRONLY, 0777)
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
