package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// пингуем урлы из списка url.txt
func main() {
	path := flag.String("file", "url.txt", "path to file") // 2 значение - дефолтное название файла

	flag.Parse()

	file, err := os.ReadFile(*path)

	if err != nil {
		panic(err.Error())
	}

	urlSlice := strings.Split(string(file), "\n") // получаем слайс урлов для пинга
	// далее логика для пингования

	respCh := make(chan int)  // канал для кодов
	errCh := make(chan error) // канал для ошибок

	for _, url := range urlSlice {
		go ping(url, respCh, errCh) // запускаем горутины для каждого урла
	}

	for range urlSlice { // знаем количество урлов (по размеру слайса), поэтому можно без waitgroup
		// // НЕВЕРНЫЙ ПОДХОД - подход ниже не подходит, тк на каждой итерации будем ждать из каждого канала, хотя по сути ответ будет только от одного
		// // 1 правильный и 1 неправильный урл, в итоге в канале с ответами будет 1 сообщение, в канале с ошибками будет 1 одно сообщение
		// // а код ниже ожидает что по каждому урлу будет и код и ошибка, поэтому будет ждать
		// errRes := <-errCh // получаем ошибки из канала
		// fmt.Println(errRes)
		// resp := <-respCh // получаем коды из канала
		// fmt.Println(resp)package main

		// ВЕРНЫЙ ПОДХОД через SELECT:
		// SELECT блокирует выполнение до тех пор, пока один из его case не будет готов к выполнению
		// SELECT автоматически выберет тот канал, в который первым придет значение.
		// Это важно, потому что для каждого URL будет либо успешный ответ, либо ошибка, но не оба значения одновременно.
		// После обработки значения из выбранного канала, цикл переходит к следующей итерации
		select {
		case err := <-errCh:
			fmt.Println(err)
		case resp := <-respCh:
			fmt.Println(resp)
			// default: // если нет ничего готового, то выполнится default
		}
	}
}

func ping(url string, respCh chan int, errCh chan error) { // для ошибок отдельный канал указываем
	res, err := http.Get(url)
	if err != nil {
		errCh <- err // записываем ошибку в канал с ошибками
		return
	}

	respCh <- res.StatusCode // записываем статус в канал с кодом
}
