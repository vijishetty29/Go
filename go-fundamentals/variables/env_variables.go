package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USER") //for windows use USERNAME
	fmt.Println("Current user logged in is ", name)
}
