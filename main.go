package main

import (
	"fmt"
	"math"
)

func main() {
	// нужно распаралеллить сложение элементов массива
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numGoroutines := 3

	sumCh := make(chan int, numGoroutines)                                            // буферизированный канал - чтобы не ждать
	var partialArrLength = int(math.Ceil(float64(len(arr)) / float64(numGoroutines))) // длина массива для каждой горутины
	arrays := splitArray(arr, partialArrLength)

	for _, partialArr := range arrays {

		go partialSum(partialArr, sumCh)
	}

	totalSum := 0

	for i := 0; i < numGoroutines; i++ { // заранее знаем сколько ждать ответов, можно за счет этого не использовать waitgroup
		totalSum += <-sumCh
	}

	fmt.Println(totalSum)

}
func splitArray(array []int, size int) [][]int {
	var result [][]int

	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		result = append(result, array[i:end])
	}

	return result
}

func partialSum(partialArr []int, ch chan int) {
	fmt.Println(partialArr)
	sum := 0
	for _, el := range partialArr {
		sum += el
	}
	ch <- sum
}
