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

func ValidAnagram(s string, t string) bool {

	s_count := make(map[rune]int)
	t_count := make(map[rune]int)

	for _, c := range s {
		s_count[c] += 1
	}

	for _, c := range t {
		t_count[c] += 1
	}

	for key, value := range s_count {
		if t_count[key] != value {
			return false
		}
	}

	return true
}
