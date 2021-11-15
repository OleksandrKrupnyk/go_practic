// преобразование данных в формате json
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
}

var JSON = `
{
"name":"Alex Fisher"
}
`

func main() {
	var p Person

	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(p)
}
