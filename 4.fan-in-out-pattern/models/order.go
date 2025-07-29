package models

import (
	"fmt"
	"time"
)

type Order struct {
	ID           int
	CustomerName string
	Items        []string
	Total        float64
	Status       string
	ProcessedAt  time.Time
}

func (o *Order) String() string {
	return fmt.Sprintf("Order{ID: %d, Customer: %s, Items: %v, Total: Rs.%.2f, Status: %s}",
		o.ID, o.CustomerName, o.Items, o.Total, o.Status)
}
