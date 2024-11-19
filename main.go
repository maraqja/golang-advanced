package main

import (
	"fmt"
	"net/http"
	"time"
)

// 10 конкурентных запросов на GET по адресу google.com
// Вывести в консоль время выполнения каждого запроса
func main() {
	t := time.Now()
	for i := 0; i < 10; i++ {
		go getHttp()
	}
	time.Sleep(time.Second*3)
	fmt.Printf("Total time: %v", time.Since(t) - time.Second*3 ) // +- 4sec без goroutine, меньее полсекунды с goroutine
	
}


func getHttp(){
	t := time.Now()
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(t)) // смотрим сколько времени прошло
	fmt.Printf("StatusCode: %v\n", resp.StatusCode)
}