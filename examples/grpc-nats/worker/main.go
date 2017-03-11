package main

import (
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"

	pb "github.com/shijuvar/gokit/examples/grpc-nats/order"
)

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	eventStore := pb.EventStore{}
	// Subscribe to subject
	natsConnection.Subscribe("Order.OrderCreated", func(msg *nats.Msg) {
		err := proto.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message %v\n", eventStore)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
