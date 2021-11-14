// Загрузка с повторениями
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	// получение сведений о текущем состоянии локального файла
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	// получение размера файла
	current := fi.Size()
	// Если локальный файл частично заполнен, указать в заголовке запроса количество загруженых байтов.
	// Отсчет начинается с 0, то есть текщий размер служит индексом следующего загружаемого байта
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}
	// http явно настроен на проверку превышения времени ожидания
	cc := &http.Client{Timeout: 5 + time.Minute}
	// Выполнение запроса для загрузки файла или части файла, если его часть уже сохранена
	res, err := cc.Do(req)

	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	// Обработка ошибочных HTTP-кодов состояний
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccessful HTTP request. Status %s"
		return fmt.Errorf(errFmt, res.Status)
	}
	// Если сервер не поддерживает обработку файлов по частям, сбросить число попыток в 0
	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}
	// копирование даных ответа в в файл
	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	return nil
}

// hasTimedOut - функция возвращает true если ошибка вызвана превышение времени ожидания
func hasTimedOut(err error) bool {
	switch err := err.(type) {
	// Ошибка url.Error может возникать из-за превышения времени ожидания в пакете net
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
		// Проверить не обнаружил ли ошибку привышения ожидания пакет net
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}

	}
	// превышение времени ожидания может быть выявлено дополнительной проверкой
	// некоторых ошибок
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

func main() {
	file, err := os.Create("ubuntu-20.04.3-desktop-amd64.iso.zsync")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	location := "http://mirror.volia.net/ubuntu-releases/focal/ubuntu-20.04.3-desktop-amd64.iso.zsync"
	err = download(location, file, 100)
	if err != nil {
		log.Println(err.Error())
	}
	fi, err := file.Stat()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Got it with %v bytes downloaded\n", fi.Size())
}
