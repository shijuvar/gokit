package main

import (
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"

	pb "github.com/shijuvar/gokit/examples/grpc-nats/order"
	"github.com/shijuvar/gokit/examples/grpc-nats/store"
)

const subject = "Order.>"

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		eventStore := pb.EventStore{}
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message in EventStore service: %+v\n", eventStore)
			store := store.EventStore{}
			store.CreateEvent(&eventStore)
			log.Println("Inserted event into Event Store")
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
