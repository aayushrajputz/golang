package main

import "fmt"

// enumarated types

type OrderStatus string

const (
	Recevied  OrderStatus = "received"
	Confirmed = "confirmed"
	Preparing = "preparing"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to", status)
}

func main() {
	changeOrderStatus()
}
