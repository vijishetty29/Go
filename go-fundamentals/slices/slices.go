package main

import (
	"fmt"
)

func main() {

	coursesSlice := make([]string, 5, 10) //length of 5 and capacity of 10
	fmt.Printf("lenght of the slice is %d and capacity of the slice is %d\n", len(coursesSlice), cap(coursesSlice))

	coursesInProg := []string{"Docker", "K8s", "Go", "Java"}

	fmt.Printf("lenght of the coursesInProg slice is %d and capacity of coursesInProg slice is %d\n", len(coursesInProg), cap(coursesInProg))

	for _, i := range coursesInProg {

		fmt.Println(i)

	}

}
