// преобразование данных в формате json в  тип interface{}
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var ks = []byte(`
{
	"firstName":"Alex",
	"lastName":"Fisher",
	"age":36,
	"education":[
		{
			"institution":"Sumy State University",
			"degree":"Bachelor of Computer Science"
		},
		{
			"institution":"Dnipro State Technical University",
			"degree":"Bachelor of Science in Mathematics"
		}
	],
	"children":[
	"Maria",
	"Denis",
	"Oksana"
	]
}
`)

func printJson(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)

	case float64:
		fmt.Println("is float", vv)
	case []interface{}:
		fmt.Println("is an array")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJson(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJson(u)
		}
	default:
		fmt.Println("Unknown type")

	}
}

func main() {
	var f interface{}

	err := json.Unmarshal(ks, &f)
	if err != nil {
		log.Fatalln(err.Error())
	}
	printJson(f)
}
