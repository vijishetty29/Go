package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	newSlice := []int{10, 20, 30}
	fmt.Println(newSlice)

	mySlice = append(mySlice, newSlice...)

	fmt.Println(mySlice)
	fmt.Printf("lenght = %d and capacity = %d\n", len(mySlice), cap(mySlice))

}
