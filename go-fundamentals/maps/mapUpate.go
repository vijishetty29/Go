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

	fmt.Println(testMap)

	testMap["B"] = 100
	fmt.Println(testMap)

	testMap["F"] = 6
	fmt.Println(testMap)

	delete(testMap, "F")
	fmt.Println(testMap)

}
