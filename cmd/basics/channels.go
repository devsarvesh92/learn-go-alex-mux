package basics

import "fmt"

func ProcessChannel(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Print("Done adding vals to channel")
}
