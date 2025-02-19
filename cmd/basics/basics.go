package arithemetic

import (
	"unicode/utf8"
)

func Add(x int, y int) int {
	return x + y
}

func GetLen(x string) int {
	return utf8.RuneCountInString(x)
}

func GetInternaStringRep(x string) []byte {
	return []byte(x)
}
