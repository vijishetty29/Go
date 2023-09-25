package main

import "fmt"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	for mapKey, mapVal := range testMap {
		fmt.Printf("We have key= %v and value= %v\n", mapKey, mapVal)
	}
}
