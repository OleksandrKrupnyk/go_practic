// Package ch04
/**
Листинг 4.2 Обработка ошибок
*/
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // Использовать только аргументы, следующие за Args[0].
	//  В первом аргументе передается имя программы

	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}
	return strings.Join(parts, " "), nil
}
