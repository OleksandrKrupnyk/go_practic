// Листинг 4.13 Восстановление после аварии
package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes() //← Вызов функции, возбуждающей аварию
}
func yikes() { //Возбуждение аварии
	panic(errors.New("something bad happened")) //← с передачей ей ошибки
}
