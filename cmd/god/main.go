package main

import (
	"fmt"
	basics "learn-go-alex-mux/cmd/basics"
)

func main() {
	fmt.Printf("Result of addition %d \n", basics.Add(2, 3))
	fmt.Printf("Number of charcaters in %d \n", basics.GetLen("Sarvesh"))
	fmt.Printf("Internal string repr %s", basics.GetInternaStringRep("Sarvesh"))
}
