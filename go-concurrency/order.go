package main

import "fmt"

type Order struct {
	ProductCode int
	Quantity    float64
	status      orderStatus
}

type InvalidOrder struct {
	order Order
	err   error
}

func (o Order) String() string {
	return fmt.Sprintf("Product code=%v, Quantity=%v, OrderStatus=%v\n",
		o.ProductCode, o.Quantity, orderStatusToText(o.status))
}

func orderStatusToText(status orderStatus) string {
	switch status {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filling:
		return "filling"
	default:
		return "Unknown status"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filling
)

var orders = []Order{}
