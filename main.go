package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 10 конкурентных запросов на GET по адресу google.com
// Вывести в консоль время выполнения каждого запроса
func main() {
	t := time.Now()

	var wg sync.WaitGroup // представляет собой по сути счетчик, который прибавляется на 1 при вызове wg.Add(1) и уменьшается на 1 при вызове wg.Done()
	
	for i := 0; i < 10; i++ {
		wg.Add(1) // говорим WaitGroup-е, что закинули 1 горутину
		// go getHttp(&wg) // обязательно передаем по указателю, тк хотим работать с WaitGroup-ой wg, а иначе в функции создастся другая waitgroup-а (копия)
		
		// 2 способ использования WaitGroup-ы, без передачи ее в функцию (функция не должна знать о WaitGroup)  
		go func() {
			getHttp()
			wg.Done()
		}()
	}
	// В результате после цикла счетчик WaitGroup-ы будет равен 10, каждая выполненная горутина когда будет выполнена уменьшает счетчик (планировщик планирует ее выполнение сам) на 1
	wg.Wait() // ждем завершения всех горутин
	fmt.Printf("Total time: %v", time.Since(t))
	
}


// func getHttp(wg *sync.WaitGroup){
func getHttp(){
	// defer wg.Done() // откладываем на после выхода из функции уменьшение счетчика WaitGroup-ы
	t := time.Now()

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(t)) // смотрим сколько времени прошло
	fmt.Printf("StatusCode: %v\n", resp.StatusCode)
}