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

	go ReceiveOrders(receivedOrdersCh)
	go ValidateOrder(receivedOrdersCh, validOrdersCh, invalidOrdersCh)
	wg.Add(1)
	go func() {
		order := <-validOrdersCh
		fmt.Println("Valid orders received ", order)
		wg.Done()
	}()
	go func() {
		order := <-invalidOrdersCh
		fmt.Println("Invalid orders received ", order.order, order.err)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(orders)
}

func ValidateOrder(in, out chan Order, errCh chan InvalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		errCh <- InvalidOrder{order: order, err: errors.New("Quantity must be greater than 0")}
	} else {
		out <- order
	}
}

func ReceiveOrders(out chan Order) {
	for _, rawOrder := range ReceiveOrder {
		var newOrder Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print("Error", err)
			continue
		}
		out <- newOrder
	}

}

var ReceiveOrder = []string{
	`{"productCode":1234,"quantity":-5,"status":2}`,
	`{"productCode":23234,"quantity":4,"status":3}`,
	`{"productCode":546234,"quantity":3,"status":4}`,
	`{"productCode":134534,"quantity":2,"status":0}`,
}
