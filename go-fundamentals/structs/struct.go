package main

import "fmt"

func main() {
	type courseMeta = struct {
		name   string
		level  string
		rating float64
	}

	gettingStarted := courseMeta{
		name:   "Docker",
		level:  "Beginner",
		rating: 4.5,
	}

	fmt.Println(gettingStarted)

	getting := new(courseMeta)
	fmt.Println(getting)

	fmt.Println("The name of the course is", gettingStarted.name)
	gettingStarted.rating = 5.0
	fmt.Println("The rating has now changed to", gettingStarted.rating)
}
