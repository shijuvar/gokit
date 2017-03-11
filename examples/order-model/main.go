package main

import (
	"fmt"
	"time"
)

func main() {
	order := &Order{
		Id: 1001,
		Customer: Customer{
			FirstName: "Alex",
			LastName:  "John",
			Email:     "alex@email.com",
			Phone:     "732-757-2923",
			Addresses: []Address{
				Address{
					Street:            "1 Mission Street",
					City:              "San Francisco",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: true,
				},
				Address{
					Street: "49 Stevenson Street",
					City:   "San Francisco",
					State:  "CA",
					Zip:    "94105",
				},
			},
		},
		Status:   "Placed",
		PlacedOn: time.Date(2016, time.April, 10, 0, 0, 0, 0, time.UTC),
		OrderItems: []OrderItem{
			OrderItem{
				Product: Product{
					Code:        "knd100",
					Name:        "Kindle Voyage",
					Description: "Kindle Voyage Wifi, 6 High-Resolution Display",
					UnitPrice:   220,
				},
				Quantity: 1,
			},
			OrderItem{
				Product: Product{
					Code:        "fint101",
					Name:        "Kindle Case",
					Description: "Fintie Kindle Voyage SmartShell Case",
					UnitPrice:   10,
				},
				Quantity: 2,
			},
		},
	}

	fmt.Println(order.ToString())
	// Change Order status
	order.ChangeStatus("Processing")
	fmt.Println("\n")
	fmt.Println(order.ToString())
}
