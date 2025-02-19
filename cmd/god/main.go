package main

import (
	"fmt"
	"learn-go-alex-mux/cmd/basics"
)

func main() {
	fmt.Printf("Result of addition %d \n", basics.Add(2, 3))
	fmt.Printf("Number of charcaters in %d \n", basics.GetLen("Sarvesh"))
	fmt.Printf("Internal string repr %v \n", basics.GetInternaStringRep("Sarvesh"))

	q, r, err := basics.Divide(4, 2)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Value of quotient %v and reminder %v", q, r)
	}
}
