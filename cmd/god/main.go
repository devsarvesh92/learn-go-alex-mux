package main

import (
	"fmt"
	"learn-go-alex-mux/cmd/basics"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Result of addition %d \n", basics.Add(2, 3))
	fmt.Printf("Number of charcaters in %d \n", basics.GetLen("Sarvesh"))
	fmt.Printf("Internal string repr %v \n", basics.GetInternaStringRep("Sarvesh"))

	fmt.Println("*******************************************")

	q, r, err := basics.Divide(4, 2)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Value of quotient %v and reminder %v \n", q, r)
	}

	fmt.Println("*******************************************")

	fmt.Printf("Find number %v \n", basics.SearchNumber([]int{3, 2, 1}, 1))

	fmt.Printf("Is Valid anagram %v \n", basics.ValidAnagram("racecar", "carrace"))
	fmt.Printf("Is Valid anagram %v \n", basics.ValidAnagram("jar", "jam"))

	fmt.Printf("Length of string %v \n", basics.GetLen("abc"))

	fmt.Println("*******************************************")

	scrapper := basics.ScrapperFactory["web"]
	result := scrapper.Scrape("abc")
	fmt.Println("Result", result.Content, result.TimeTaken)

	fmt.Println("*******************************************")

	numbers := []int{1, 2, 3, 4}
	fmt.Printf("Memory location: %p\n", &numbers)
	basics.GetSquares(&numbers)
	fmt.Println("Result", numbers)
	fmt.Printf("Memory location result: %p\n", &numbers)
	fmt.Println("*******************************************")

	fmt.Println("*******************************************")

	var wg sync.WaitGroup
	var mu sync.RWMutex
	var nums []int = []int{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			nums = append(nums, i)
			mu.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			fmt.Println(nums)
			mu.RUnlock()
		}()
	}

	wg.Wait()

	fmt.Println(nums)
	fmt.Println("*******************************************")

	channel := make(chan int, 10)
	go basics.ProcessChannel(channel)

	for i := range channel {
		time.Sleep(time.Second)
		fmt.Print(i)
	}
	fmt.Print("Done processing records")

	fmt.Println("*******************************************")

	var fwg sync.WaitGroup
	chicken_channel := make(chan string)
	tofu_channel := make(chan string)

	fwg.Add(2)
	go func() {
		defer fwg.Done()
		basics.CheckChickenSales(chicken_channel, []string{"abc.com", "costco.com", "kfc.com"})
	}()
	go func() {
		defer fwg.Done()
		basics.CheckTofuSales(tofu_channel, []string{"abc.com", "costco.com", "kfc.com"})
	}()
	go func() {
		defer fwg.Done()
		basics.SendMessage(chicken_channel, tofu_channel)
	}()

	fwg.Wait()

	fmt.Println("*******************************************")
	fmt.Println("Result of number addition", basics.SumSlice([]int{1, 2, 3, 4}))
}
