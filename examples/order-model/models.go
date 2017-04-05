package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street, City, State, Zip string
	IsShippingAddress        bool
}

type Customer struct {
	FirstName, LastName, Email, Phone string
	Addresses                         []Address
}

func (c Customer) ToString() string {
	return fmt.Sprintf("Customer: %s %s, Email:%s", c.FirstName, c.LastName, c.Email)
}
func (c Customer) ShippingAddress() string {
	for _, v := range c.Addresses {
		if v.IsShippingAddress == true {
			return fmt.Sprintf("%s, %s, %s, Zip - %s", v.Street, v.City, v.State, v.Zip)
		}
	}
	return ""
}

type Order struct {
	Id int
	Customer
	PlacedOn   time.Time
	Status     string
	OrderItems []OrderItem
}

func (o *Order) GrandTotal() float64 {
	var total float64
	for _, v := range o.OrderItems {
		total += v.Total()
	}
	return total
}
func (o *Order) ToString() string {
	var orderStr string
	orderStr = fmt.Sprintf("Order#:%d, OrderDate:%s, Status:%s, Grand Total:%f\n", o.Id, o.PlacedOn, o.Status, o.GrandTotal())
	orderStr += o.Customer.ToString()
	orderStr += fmt.Sprintf("\nOrder Items:")
	for _, v := range o.OrderItems {
		orderStr += fmt.Sprintf("\n")
		orderStr += v.ToString()
	}
	orderStr += fmt.Sprintf("\nShipping Address:")
	orderStr += o.Customer.ShippingAddress()
	return orderStr
}
func (o *Order) ChangeStatus(newStatus string) {
	o.Status = newStatus
}

type OrderItem struct {
	Product
	Quantity int
}

func (item OrderItem) Total() float64 {
	return float64(item.Quantity) * item.UnitPrice
}
func (item OrderItem) ToString() string {
	itemStr := fmt.Sprintf("Code:%s, Product:%s -- %s, UnitPrice:%f, Quantity:%d, Total:%f",
		item.Code, item.Name, item.Description, item.UnitPrice, item.Quantity, item.Total())
	return itemStr

}

type Product struct {
	Code, Name, Description string
	UnitPrice               float64
}
