package main

import (
	"fmt"
)

func main() {
	//courseGoLength := 240     // GO length
	//courseDockerLength := 240 // Docker Course Length
	//We declare values inline of the if statement only available locally within the if conditional statements only

	if courseGoLength, courseDockerLength := 240, 120; courseGoLength > courseDockerLength {
		fmt.Println("Go Course takes longer time than Docker")
		if courseGoLength > 200 {
			fmt.Println("This can put viewers to sleep!")
		}
	} else if courseGoLength < courseDockerLength {
		fmt.Println("Go Course takes lesser time than Docker")
	} else {
		fmt.Println("Go Course and Docker take equal time")
	}
}
