// Листинг 4.22 safely.Go
package safely

import (
	"log"
)

type GoDoer func() //← GoDoer – простая функция без параметров
// вызывает функцию
// как сопрограмму
// и обрабатывает все аварии
func Go(todo GoDoer) {
	go func() { //← Сначала запускается анонимная функция
		defer func() {
			if err := recover(); err != nil { // Анонимная функция обрабатывает аварии по обычному сценарию восстановления
				log.Printf("Panic in safely.Go: %s", err)
			}
		}()
		todo() //← Затем эта функция вызывает переданную в параметре функцию GoDoer
	}()
}
