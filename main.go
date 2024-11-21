package main

import (
	"fmt"
	"net/http"
	"sync"
)

// За счет WG контролим, что определенное нужное нам кол-во горитун выполнилось
// За счет каналов обеспечиваем коммуникацию горутин (вложенная функция с вызывающей)
func main() {

	var wg sync.WaitGroup
	code := make(chan string) // канал для передачи кода ответа

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttp(code, i)
			wg.Done()
		}()
	}

	// В рамках этой горутины ждем, когда все горутины завершатся (когда счетчик waitgroup станет 0) и закрываем канал для завершения горутины main
	go func() {
		wg.Wait()
		close(code) // Закрываем канал, что приведет к завершению range цикла ниже
	}()

	// В рамках этого цикла бесконечно получаем новые значения из канала
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
