package main

import (
	"fmt"
)

func main() {

	coursesInProg := []string{"Docker", "K8s", "Go", "Java"}

	coursesCompleted := []string{"Docker", "Go", "Java"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("There is a clash on ", i)
			}
		}

	}

}
