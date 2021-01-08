/*
List 3.1
*/
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// Вызов функции echo как go-подпрограммы
	go echo(os.Stdin, os.Stdout)
	// 30-секундная пауза
	time.Sleep(30 * time.Second)
	//  Вывод сообщения о завершении ожидания
	fmt.Printf("\nTime out.\n")
	// Выход из программы. При этом сопрограмма будет остановлена.
	os.Exit(0)
}

/*
Функция echo является обычной функцией
*/
func echo(in io.Reader, out io.Writer) {
	// io.Copy скопирует данные из os.Reader в os.Writer
	io.Copy(out, in)
}
