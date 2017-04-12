package main

import (
	"log"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"

	pb "github.com/shijuvar/gokit/examples/grpc-nats/order"
)

var orderServiceUri string

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file not found")
	}
	orderServiceUri = viper.GetString("discovery.orderservice")
}
func main() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	natsConnection.Subscribe("Discovery.OrderService", func(m *nats.Msg) {
		orderServiceDiscovery := pb.ServiceDiscovery{OrderServiceUri: orderServiceUri}
		data, err := proto.Marshal(&orderServiceDiscovery)
		if err == nil {
			natsConnection.Publish(m.Reply, data)
		}
	})
	// Keep the connection alive
	runtime.Goexit()
}
