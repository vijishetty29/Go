package main

import "fmt"

func main() {
	const (
		FIZZ = 3
		BUZZ = 5
	)

	for i := 0; i < 100; i++ {

		p := false
		if i%FIZZ == 0 {
			fmt.Printf("FIZZ")
			p = true
		}
		if i%BUZZ == 0 {
			fmt.Printf("BUZZ")
			p = true
		}

		if !p {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}

}
