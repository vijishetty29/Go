package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// 	i := 0
	// loop:
	// 	if i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 		goto loop
	// 	}

	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i
	}

	fmt.Printf("%v", array)
}
