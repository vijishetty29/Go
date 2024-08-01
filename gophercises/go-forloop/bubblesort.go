package main

import "fmt"

func main() {
	array := []int{5, -1, 4, 7, 2}
	fmt.Println("Unsorted array %v", array)
	bubblesort(array)
	fmt.Println("Sorted array %v", array)
}

func bubblesort(array []int) {

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

}
