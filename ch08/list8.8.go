// преобразование http-ответа в ошибку
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Error - структура для хранения данных об ошибке
type Error struct {
	HTTPcode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

// метод реализует интерфейс Error
func (e Error) Error() string {
	fs := "HTTP : %d, Code: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPcode, e.Code, e.Message)
}

func get(u string) (*http.Response, error) {
	// использование get для получения ресурса и возврата любых ошибок
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}
	// Код ошибки
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// провертка типа соадержимого и возврат ошибки, если он не правильный
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Unknown error HTTP status %s"
			return res, fmt.Errorf(sm, res.Status)
		}
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		// Разобрать JSON ответ записать данные в структкру и вернуть ошибку
		var data struct {
			Err Error `json:"error"`
		}
		// Разобрать JSON в структуру data
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := " Unable to parse json: %s, HTTP Status %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		// Добавление кода http состояния в екземпляр error
		data.Err.HTTPcode = res.StatusCode
		// возврат польовательской ошибки с ответом
		return res, data.Err
	}
	// при отсутсвии ошибки возвращать обычный ответ
	return res, nil

}

func main() {
	res, err := get("http://localhost:8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("%s", b)
}
