// Листинг 4.8 Обработка двух различных ошибок
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("the request timed out")     // Создание экземпляра ошибки превышения времени ожидания
var ErrRejected = errors.New("the request was rejected") // Создание экземпляра ошибки отказа

var random = rand.New(rand.NewSource(35)) // Генерация случайных чисел с помощью  фиксированного источника
func main() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.(", ErrTimeout, ")")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
