package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "viji shetty"
	course := "getting started with go course"

	fmt.Println(converter(name, course))
}

func converter(name, course string) (n, c string) {
	name = strings.ToUpper(name)
	course = strings.Title(course)
	return name, course
}
