package basics

func GetSquares(input *[]int) []int {
	val := *input
	for x := range val {
		val[x] = val[x] * val[x]
	}
	return val
}
