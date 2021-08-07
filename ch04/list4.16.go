// Листинг 4.16 Очистка
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser

	file, err := OpenCSV("data.csv") // Вызов функции OpenCSV и
	// обработка всех ошибок. Эта реализация всегда возвращает ошибку
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer func(file io.ReadCloser) {
		err := file.Close()
		if err != nil {

		}
	}(file) // Использовать отложенную функцию,
	// чтобы гарантировать закрытие файла при любых обстоятельствах
	// Что-то делается с файлом. Обычно здесь помещается код для работы с файлом
}
func OpenCSV(filename string) (file *os.File, err error) {
	defer func() { // Отложенная обработка возникшей ошибки с передачей ее в функцию main
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()

	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}
	RemoveEmptyLines(file) //← Вызов функции RemoveEmptyLines,
	return file, err       // постоянно вызывающей аварии
}
func RemoveEmptyLines(f *os.File) {
	panic(errors.New("failed parse")) //← Вместо удаления пустых строк всегда возбуждает аварию
}
