package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Error - Тип для хранения информации об ошибке
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func JSONError(w http.ResponseWriter, e Error) {
	// Обертывание структуры Error анонимной структурой со свойством Error
	data := struct {
		Err Error `json:"error"`
	}{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	// установка MIME ответа в заголовке
	w.Header().Set("Content-Type", "application/json")
	// Гарантированное правильное установление кода ошибки
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     123,
		Message:  "An Error Occurred",
	}
	JSONError(w, e)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
