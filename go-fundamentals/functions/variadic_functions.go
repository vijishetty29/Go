package main

import (
	"fmt"
)

func main() {
	bestFinish := championFinishes(12, 5, 4, 7, 3, 5, 2, 8, 5)

	fmt.Println("Best finish is ", bestFinish)
}

func championFinishes(finishes ...int) int {
	bestFinish := finishes[0]
	for _, i := range finishes {
		if i < bestFinish {
			bestFinish = i
		}
	}
	return bestFinish
}
