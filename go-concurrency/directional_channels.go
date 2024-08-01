package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersCh = make(chan Order)
	var validOrdersCh = make(chan Order)
	var invalidOrdersCh = make(chan InvalidOrder)

	go receiveOrders(receivedOrdersCh)
	go validateOrder(receivedOrdersCh, validOrdersCh, invalidOrdersCh)
	wg.Add(1)
	go func(validOrdersCh <-chan Order) {
		order := <-validOrdersCh
		fmt.Println("Valid orders received ", order)
		wg.Done()
	}(validOrdersCh)
	go func(invalidOrdersCh <-chan InvalidOrder) {
		order := <-invalidOrdersCh
		fmt.Println("Invalid orders received ", order.order, order.err)
		wg.Done()
	}(invalidOrdersCh)

	wg.Wait()
	fmt.Println(orders)
}
