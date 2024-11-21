package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	code := make(chan int) // канал для передачи кода ответа

	go getHttp(code) // передаем канал в горутину

	statusCode := <-code //  Читаем из канала в переменную
	// функция main блокируется до появления в канале хотя бы одного значения (тк функция main запускается в горутине)
	fmt.Printf("StatusCode: %v\n", statusCode)
}

func getHttp(codeCh chan int) { // принимаем канал для записи

	t := time.Now()

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(t)) // смотрим сколько времени прошло

	codeCh <- resp.StatusCode // кладем в канал StatusCode
}
