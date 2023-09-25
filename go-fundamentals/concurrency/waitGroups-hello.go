package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2) // Step 1. Add waitgroups
	//number of go rountines that many should be included
	//so that the main programs knows to exit when all wait groups are marked done.

	go func() { //Step 2. Add go keyword
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("World")
	}()

	waitGrp.Wait() //Step 3. Wait
	//Without this wait the main function will exit
	// without waiting for the go routines to finish the task
}
