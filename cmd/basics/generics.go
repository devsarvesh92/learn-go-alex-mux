package basics

func SumSlice[T int | float32](slice []T) T {
	var result T
	for _, i := range slice {
		result += i
	}
	return result
}
