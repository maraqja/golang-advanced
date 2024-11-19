package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello from goroutine") // заставляем выполняться в отдельной горутине (планируется планировщиком выполнение)
	fmt.Println("Hello from main")
	go fmt.Println("Hello from goroutine 2") // будет разный порядок выполнеия горутин
	time.Sleep(time.Nanosecond) // ожидаем завершения горутины (успеет все выполниться за наносекунду)
}