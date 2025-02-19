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
		fmt.Printf("Value of quotient %v and reminder %v \n", q, r)
	}

	fmt.Printf("Find number %v \n", basics.SearchNumber([]int{3, 2, 1}, 1))

	fmt.Printf("Is Valid anagram %v \n", basics.ValidAnagram("racecar", "carrace"))
	fmt.Printf("Is Valid anagram %v \n", basics.ValidAnagram("jar", "jam"))

	fmt.Printf("Length of string %v \n", basics.GetLen("abc"))

}
