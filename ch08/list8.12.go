// регистрация пути к API с включение версии
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type testMessage struct {
	Message string `json:"message"`
}

func displayError(w http.ResponseWriter, r *http.Request) {
	data := testMessage{"A test message"}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "applicant/json")
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/api/v1/test", displayError)
	http.ListenAndServe(":8080", nil)
}
