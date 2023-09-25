package main

import (
	"fmt"
	"time"
)

func main() {

	for timer := 5; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}

		if timer%2 == 0 {
			continue
		}

		fmt.Println(timer)
		time.Sleep(1)
	}

}
