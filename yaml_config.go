/*
List 2.10
*/
package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {

	var file *yaml.File
	var err error

	file, err = yaml.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	path, _ := file.Get("path")
	enabled, _ := file.Get("enabled")
	fmt.Println(path)
	fmt.Println(enabled)
}
