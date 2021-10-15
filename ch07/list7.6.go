// загрузка статических файлов в память и их обслуживание
package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// структура даних для хранения файла в памяти
type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile
var mutex = new(sync.RWMutex)

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	// Загрузка файла из кеша если он у него там есть
	v, found := cache[req.URL.Path]
	mutex.RUnlock()

	if !found {
		mutex.Lock()
		defer mutex.Unlock()
		fileName := "./ch07/files/" + req.URL.Path
		f, err := os.Open(fileName)
		defer f.Close()
		// Обработка ошибок открытия файла
		if err != nil {
			http.NotFound(res, req)
			log.Println(err.Error())
			return
		}
		// Перемення по буфер (область памяти)
		var b bytes.Buffer
		// Копирование содержимого из файла в буфер памяти
		// Обработка ошибок копирования
		if _, err := io.Copy(&b, f); err != nil {
			http.NotFound(res, req)
			log.Println(err.Error())
			return
		}
		// Поместить прочитаные байти в Reader для
		// дальнейшего использования
		r := bytes.NewReader(b.Bytes())
		// Информация о файле
		info, _ := f.Stat()

		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}

func main() {
	// создание карты файлов
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
