/*
List 2.9
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
	Parts   []int64
}

func main() {
	var file *os.File
	var decoder *json.Decoder
	var conf *configuration
	var err error

	conf = &configuration{}
	// Чтение файла
	file, _ = os.Open("config.json")
	defer file.Close()
	// Декодер
	decoder = json.NewDecoder(file)
	// Процесс декодирования
	err = decoder.Decode(conf)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(conf.Path)
	fmt.Println(*conf)
}
