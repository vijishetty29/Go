package main

import "fmt"

func main() {
	titles := make(map[string]int)
	titles["Shetty"] = 7
	titles["Nino"] = 5

	recentTitles := map[string]int{
		"Viji": 6,
		"Mani": 5,
	}

	fmt.Println(titles)
	fmt.Println(recentTitles)
}
