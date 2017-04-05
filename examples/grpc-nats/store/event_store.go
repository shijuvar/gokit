package store

import (
	pb "github.com/shijuvar/gokit/examples/grpc-nats/order"
)

// EventStore provides CRUD operations against the collection "orders"
type EventStore struct {
}

// CreateEvent inserts the value of struct Order into collection.
func (store EventStore) CreateEvent(order *pb.EventStore) error {
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("natsdemo").C("events")
	err := col.Insert(order)
	return err
}

// GetEvents returns all documents from the collection.
func (store OrderStore) GetEvents() []pb.EventStore {
	var events []pb.EventStore
	session := mgoSession.Copy()
	defer session.Close()
	col := session.DB("natsdemo").C("events")
	iter := col.Find(nil).Iter()
	result := pb.EventStore{}
	for iter.Next(&result) {
		events = append(events, result)
	}
	return events
}
