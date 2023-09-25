package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name   = "Vij Shetty"
	course = "Go Course"
	module = "2" // Uh no it is a string
	clip   = 4
)

func main() {
	fmt.Println("Name is ", name, " and Course ", course, ".")
	fmt.Println("Module is ", module, " and Clip ", clip, ".")
	fmt.Println("Name is of type ", reflect.TypeOf(name))
	fmt.Println("Module is of type ", reflect.TypeOf(module))

	moduleInt, err := strconv.Atoi(module)

	if err == nil {
		total := moduleInt + clip
		fmt.Println("Total is ", total)
	}
	fmt.Println("Memory address of course variable is ", &course)

	var ptr *string = &course
	fmt.Println("Pointer address for course is ", ptr, " Value of the course is ", *ptr)
}
