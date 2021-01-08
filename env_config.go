/*
List 2.12
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
Получение данных с переменных окружения
*/
func main() {
	var port string
	port = os.Getenv("PORT")
	fmt.Println("Port is :", port, " номер")
	http.HandleFunc("/", homePage)
	http.ListenAndServe("127.0.0.1:"+port, nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
