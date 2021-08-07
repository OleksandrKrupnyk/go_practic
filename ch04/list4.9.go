// Листинг 4.9 Ошибки и аварии
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Divide 2 by 0")
	divide(2, 0)
}
func precheckDivide(a, b int) (int, error) { //← Функция precheckDivide возвращает
	if b == 0 { //   ошибку при делении на 0
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}
func divide(a, b int) int { //← Обычная функция divide просто выполняет
	return a / b //   операцию деления без проверок
}
