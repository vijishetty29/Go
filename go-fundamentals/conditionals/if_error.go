package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("./test1.txt")

	if err != nil {
		fmt.Println("An error has occured ", err)
	}

	fmt.Println("An error has not occured ", err)

}
