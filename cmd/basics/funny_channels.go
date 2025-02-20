package basics

import (
	"fmt"
	"time"
)

func CheckChickenSales(c chan string, websites []string) {
	for _, website := range websites {
		time.Sleep(time.Second)
		fmt.Println("Sale found on", website)
		c <- website
		break
	}
	close(c)
}

func CheckTofuSales(c chan string, websites []string) {
	for _, website := range websites {
		time.Sleep(time.Second)
		fmt.Println("Tofu sale found on", website)
		c <- website
		break
	}
}

func SendMessage(chicken chan string, tofu chan string) {
	select {
	case website := <-chicken:
		fmt.Println("Found deal on chicken", website)
	case website := <-tofu:
		fmt.Println("Found deal on tofu", website)
	}
}
