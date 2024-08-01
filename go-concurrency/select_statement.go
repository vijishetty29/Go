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

	go ReceiveOrders(receivedOrdersCh)
	go ValidateOrder(receivedOrdersCh, validOrdersCh, invalidOrdersCh)
	wg.Add(1)

	go func(validOrdersCh <-chan Order, invalidOrdersCh <-chan InvalidOrder) {

		select {
		case order := <-validOrdersCh:
			fmt.Println("Valid orders received ", order)
		case order := <-invalidOrdersCh:
			fmt.Println("Invalid orders received ", order.order, order.err)
		default:
			fmt.Println("defaut case")
		}

		wg.Done()
	}(validOrdersCh, invalidOrdersCh)

	wg.Wait()
	fmt.Println(orders)
}
