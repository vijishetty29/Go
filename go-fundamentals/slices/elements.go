package main

import "fmt"

func main() {
	slices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slices)

	fmt.Println(slices[3])

	slices[2] = 0

	fmt.Println(slices)

	sliceOfSlices := slices[3:6]

	fmt.Println(sliceOfSlices)

	sliceOfSlices1 := slices[:6]

	fmt.Println(sliceOfSlices1)

	sliceOfSlices2 := slices[7:]

	fmt.Println(sliceOfSlices2)
}
