package basics

import (
	"fmt"
	"strings"
)

func GetLen(s string) int {
	s_rune := []rune(s)
	var sb strings.Builder
	for _, x := range s_rune {
		sb.WriteRune(x)
	}
	fmt.Printf("Val %v \n", sb.String())
	return len(s_rune)
}
