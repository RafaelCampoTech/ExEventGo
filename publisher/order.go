package main

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusCompleted OrderStatus = "Completed"
	OrderStatusCancelled OrderStatus = "Cancelled"
)

type OrderCreated struct {
	OrderID        string
	OrderDetails   string
	OrderCreatedAt time.Time
	OrderQty       int
	OrderStatus    OrderStatus
}
