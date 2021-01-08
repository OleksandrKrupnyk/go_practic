/*
List 2.14
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
Сервер Антипатерн
*/
func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homePage2)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

/*
Завершение работы сервера
*/
func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

/*
Строка
*/
func homePage2(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "This is home page.")
}
