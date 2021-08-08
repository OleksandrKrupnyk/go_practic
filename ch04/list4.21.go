// Листинг 4.23 Использование safely.Go для перехвата аварий
package main

import (
	"errors"
	"github.com/Masterminds/cookoo/safely" //← Импорт пакета safely
	"time"
)

func message() { //← Определение функции обратного вызова, соответствующей типу GoDoer
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message) // ← Заменяет go message
	println("Outside goroutine")
	time.Sleep(1000) // ← Гарантирует выполнение сопрограммы
} //перед завершением работы программы
