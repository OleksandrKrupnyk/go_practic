/**
List 1.4
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error 1")
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error 2")
		panic(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
