package main

import (
	"fmt"
)

func main() {
	name := "Viji Shetty"
	course := "Go Course"

	fmt.Println("Hi, ", name, " Value of the course is ", course)

	updateCourse(&course)

	fmt.Println("Currently watching course value ", course)
}

func updateCourse(course *string) string {
	*course = "Getting started with Docker"
	fmt.Println("Updated course value is ", *course)
	return *course
}
