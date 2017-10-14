package main

import (
	"context"
	"log"
	"net"

	stan "github.com/nats-io/go-nats-streaming"
	"google.golang.org/grpc"

	"github.com/shijuvar/gokit/examples/nats-streaming/pb"
	"github.com/shijuvar/gokit/examples/nats-streaming/store"
)

const (
	port      = ":50051"
	clusterID = "test-cluster"
	clientID  = "event-store"
)

type server struct{}

// CreateOrder RPC creates a new Event into EventStore
func (s *server) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {
	// Persist data into EventStore database
	command := store.EventStore{}
	// Persist events as immutable logs into CockroachDB
	err := command.CreateEvent(in)
	if err != nil {
		return nil, err
	}
	// Publish event on NATS Streaming Server
	go publishEvent(in)
	return &pb.Response{IsSuccess: true}, nil
}

// GetEvents RPC gets events from EventStore by given AggregateId
func (s *server) GetEvents(ctx context.Context, in *pb.EventFilter) (*pb.EventResponse, error) {
	eventStore := store.EventStore{}
	events := eventStore.GetEvents(in)
	return &pb.EventResponse{Events: events}, nil
}

// publishEvent publish an event via NATS Streaming server
func publishEvent(event *pb.Event) {
	// Connect to NATS Streaming server
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer sc.Close()
	channel := event.Channel
	eventMsg := []byte(event.EventData)
	// Publish message on subject (channel)
	sc.Publish(channel, eventMsg)
	log.Println("Published message on channel: " + channel)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterEventStoreServer(s, &server{})
	s.Serve(lis)
}
