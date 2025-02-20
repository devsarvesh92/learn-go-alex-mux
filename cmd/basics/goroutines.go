package basics

import (
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println("Number", number)
}

func SafeWorker(numbers *[]int, number int) {
	*numbers = append(*numbers, number)
}
