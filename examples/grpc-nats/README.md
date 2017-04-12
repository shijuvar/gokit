## Components in the Demo App
* order - Protocol Buffers definition file.
* client - A gRPC client app, which is used for demonstrate Request-Reply messaging of NATS by sending a request on a subject.
* discovery - A simple service discovery app, which is used for demonstrate Request-Reply messaging of NATS by publishing a reply for the request on a subject.
* server - A gRPC server app, which is used for demonstrate Publish-Subscribe messaging of NATS by publishing messages.
* eventstore - A NATS client app that subscribes messages by subscribing messages on a subject wildcard.
* worker1 - A NATS client app that subscribes messages via subscriber queue group.
* worker2 - A NATS client app that subscribes messages via subscriber queue group.
* store - Persistence layer that performs the persistence operations on MongoDB.

## Compile Proto files
Run the command below from the grpc-nats directory:
protoc -I order/ order/order.proto --go_out=plugins=grpc:order

## Technologies Used: 
* NATS
* gRPC
* MongoDB
