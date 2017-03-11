## To compile proto files
Run the command below from the grpc-nats directory:
protoc -I order/ order/order.proto --go_out=plugins=grpc:order
