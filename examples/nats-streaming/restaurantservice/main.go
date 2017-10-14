package main

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	stan "github.com/nats-io/go-nats-streaming"

	"github.com/shijuvar/gokit/examples/nats-streaming/pb"
)

const (
	clusterID = "test-cluster"
	clientID  = "restaurant-service"
	channel   = "order-notification"
	durableID = "restaurant-service-durable"
)

func main() {
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)

	if err != nil {
		log.Fatal(err)
	}
	// Subscribe with manual ack mode, and set AckWait to 60 seconds
	aw, _ := time.ParseDuration("60s")
	sc.Subscribe(channel, func(msg *stan.Msg) {
		msg.Ack() // Manual ACK
		order := pb.Order{}
		// Unmarshal JSON that represents the Order data
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Print(err)
			return
		}
		// Handle the message
		log.Printf("Subscribed message from clientID - %s for Order: %+v\n", clientID, order)

	}, stan.DurableName(durableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
	runtime.Goexit()
}
