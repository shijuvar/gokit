package store

import (
	pb "github.com/shijuvar/gokit/examples/grpc-nats/order"
)

// OrderStore provides CRUD operations against the collection "orders"
type OrderStore struct {
}

// CreateOrder inserts the value of struct Order into collection.
func (store OrderStore) CreateOrder(order *pb.Order) error {
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("natsdemo").C("orders")
	err := col.Insert(order)
	return err
}

// GetOrders returns all documents from the collection.
func (store OrderStore) GetOrders() []pb.Order {
	var orders []pb.Order
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("natsdemo").C("orders")
	iter := col.Find(nil).Iter()
	result := pb.Order{}
	for iter.Next(&result) {
		orders = append(orders, result)
	}
	return orders
}
