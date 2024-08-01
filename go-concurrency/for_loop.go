package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersCh = make(chan Order)
	var validOrdersCh = make(chan Order)
	var invalidOrdersCh = make(chan InvalidOrder)

	go ReceiveOrdersClose(receivedOrdersCh)
	go ValidateOrders(receivedOrdersCh, validOrdersCh, invalidOrdersCh)
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
func ValidateOrders(in, out chan Order, errCh chan InvalidOrder) {
	//order := <-in
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- InvalidOrder{order: order, err: errors.New("Quantity must be greater than 0")}
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
}
func ReceiveOrdersClose(out chan Order) {
	for _, rawOrder := range ReceiveOrder {
		var newOrder Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print("Error", err)
			continue
		}
		out <- newOrder
	}
	close(out)
}
