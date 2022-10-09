package entity

import "time"

type Order struct {
	OrderId      int
	CustomerName string
	OrderedAt    time.Time
}
