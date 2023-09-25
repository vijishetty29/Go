package main

import (
	"fmt"
)

func main() {

	switch "Docker" {
	case "docker":
		fmt.Println("Case 1 : The dockers \"d\" is small in this case")
	case "Docker":
		fmt.Println("Case 2 : The dockers \"d\" is capital in this case")
	case "D":
		fmt.Println("Case 3 : The Docker is only D")
	default:
		fmt.Println("No case found!")
	}
}
