package main

func NumInList(numbers []int, n int) bool {
	for _, num := range numbers {
		if num == n {
			return true
		}
	}
	return false
}
