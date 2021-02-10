/*
 List 3.6   Подсчет слов с блокировками
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Структура для хранения хранения сло и количества их вхождения в файле
type words struct {
	sync.Mutex // Структура words теперь наследует мьютекс
	found      map[string]int
}

// создание нового экземпляра слова
func newWords() *words {
	return &words{found: map[string]int{}}
}

/**
  Фиксируем количество вхождений n слова word
*/
func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	// Если не найдено то создать
	if !ok {
		w.found[word] = n
		return
	}
	// иначе добавить к существующему
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

func main() {
	var wg sync.WaitGroup
	w := newWords()
	for _, f := range os.Args[1:] {
		fmt.Println(f)
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:") // Перед завершением
	for word, count := range w.found {               //программы вывести
		if count > 5 { //результаты поиска
			fmt.Printf("%s: %d\n", word, count)
		}
	}

}
