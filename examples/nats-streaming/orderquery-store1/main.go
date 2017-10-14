package main

import (
	"encoding/json"
	"log"
	"runtime"

	stan "github.com/nats-io/go-nats-streaming"

	"github.com/shijuvar/gokit/examples/nats-streaming/pb"
	"github.com/shijuvar/gokit/examples/nats-streaming/store"
)

const (
	clusterID  = "test-cluster"
	clientID   = "order-query-store1"
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
			queryStore := store.QueryStore{}
			// Perform data replication for query model into CockroachDB
			err := queryStore.SyncOrderQueryModel(order)
			if err != nil {
				log.Printf("Error while replicating the query model %+v", err)
			}
		}
	}, stan.DurableName(durableID),
	)
	runtime.Goexit()
}
