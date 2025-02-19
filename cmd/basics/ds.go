package basics

import "sort"

func SearchNumber(numbers []int, number int) bool {
	start, end := 0, len(numbers)-1
	sort.Ints(numbers)

	for start <= end {
		middle := (start + end) / 2
		if number == numbers[middle] {
			return true
		} else if number < numbers[middle] {
			end = middle - 1
		} else if number > numbers[middle] {
			start = middle + 1
		}

	}
	return false
}
