package main

import (
	"fmt"
	"net/http"
)

func main() {
	code := make(chan string) // канал для передачи кода ответа

	for i := 0; i < 10; i++ {
		go getHttp(code, i)
	}

	for res := range code {
		fmt.Printf("StatusCode: %v\n", res) // получаем все коды из канала (они там будут в рандом порядке, тк какя раньше выполнилась, та раньше и записала результат)
	}

}

func getHttp(codeCh chan string, iteration int) { // принимаем канал для записи

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	codeCh <- fmt.Sprintf("[%d] %d", iteration, resp.StatusCode) // кладем в канал StatusCode
}
