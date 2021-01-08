/*
List 3.3
*/
package main

import (
	"compress/gzip"
	"io"
	"os"
)

/*

 */
func main() {
	for _, file := range os.Args[1:] {
		compress(file)
	}

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
