// Листинг 5.2 Журналирование в файл
package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.OpenFile("./file.log", os.O_RDWR|os.O_APPEND, 0666) //← Создание файла журнала
	defer logfile.Close()                                                //← Гарантировать закрытие
	logger := log.New(logfile, "info ", log.LstdFlags|log.Lshortfile)
	logger.Println("This is a regular message.") //   Отсылка сообщений
	logger.Fatalln("This is a fatal error.")
	logger.Println("This is the end of the function.") //← Как и раньше, никогда
}
