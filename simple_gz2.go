/*
List 3.4  Параллельное сжатие файлов с ожиданием завершения
группы
*/
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

/*

 */
func main() {
	var wg sync.WaitGroup // Нет необходимости инициализировать WaitGroup
	var i int = -1        // Так как переменная i необходима за пределами цикла
	var file string       // она объявляется именно здесь
	for i, file = range os.Args[1:] {
		wg.Add(1) // Для каждого файла сообщить группе что ожидается
		// выполнение еще одной операции сжатия
		go func(filename string, number int) { // Эта функция вызывает функцию сжатия
			compress(filename) // и уведомляет группу ожидания о ее завершении
			fmt.Printf("%02d %s\n", number+1, filename)
			wg.Done()
		}(file, i) // Поскольку вызов сопрограммы происходит в цикле for
		// требуется небольшая хитрость, чтобы передать параметр
	}
	/*
		Внешняя сопрограмма (main) ожидает  пока все сопрограммы, выполняющие
		сжатие, вызовут wg.Done
	*/
	wg.Wait()
	fmt.Printf("Compressed %d files\n", i+1)

}

/*

 */
func compress(fileName string) error {
	// Открыть исходный файл для чтения
	in, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer in.Close()
	//  Открыть файл архива с расширением .gz и именем исходного файла
	out, err := os.Create(fileName + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()
	// Сжать данные и записать в соответствующий файл с помощью  gzip.Writer
	// Функция io.Copy выполняет
	//  необходимое копирование
	gzout, _ := gzip.NewWriterLevel(out, gzip.BestCompression)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
