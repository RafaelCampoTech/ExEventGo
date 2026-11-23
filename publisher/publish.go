package main

import (
	"fmt"
)

func PublishOrder(order OrderCreated) {
	fmt.Print(order.OrderDetails)
}
