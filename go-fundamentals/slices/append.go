package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 1, 4) //length of 5 and capacity of 10
	fmt.Printf("START Length of the slice is %d and capacity of the slice is %d\n", len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("lenght = %d and capacity = %d\n", len(mySlice), cap(mySlice))

	}

}
