package main

import (
	"encoding/json"
	"log"
	"runtime"

	stan "github.com/nats-io/go-nats-streaming"

	"github.com/shijuvar/gokit/examples/nats-streaming/pb"
)

const (
	clusterID  = "test-cluster"
	clientID   = "order-query-store2"
	channel    = "order-notification"
	durableID  = "store-durable"
	queueGroup = "order-query-store-group"
)

func main() {
	// Connect to NATS Streaming server
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)

	if err != nil {
		log.Fatal(err)
	}
	sc.QueueSubscribe(channel, queueGroup, func(msg *stan.Msg) {
		order := pb.Order{}
		err := json.Unmarshal(msg.Data, &order)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message from clientID - %s: %+v\n", clientID, order)
		}
	}, stan.DurableName(durableID),
	)
	runtime.Goexit()
}
