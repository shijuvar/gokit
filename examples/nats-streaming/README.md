## Technologies Used: 
* NATS Streaming
* gRPC
* CockroachDB

## Components in the Demo App
 *Protocol Buffers definition file.
* client - A gRPC client app, which is used for demonstrate Request-Reply messaging of NATS by sending a request on a subject.
* discovery - A simple service discovery app, which is used for demonstrate Request-Reply messaging of NATS by publishing a reply for the request on a subject.
* server - A gRPC server app, which is used for demonstrate Publish-Subscribe messaging of NATS by publishing messages.
* eventstore - A NATS client app that subscribes messages by subscribing messages on a subject wildcard.
* worker1 - A NATS client app that subscribes messages via subscriber queue group.
* worker2 - A NATS client app that subscribes messages via subscriber queue group.
* store - Persistence layer that performs the persistence operations on MongoDB.

## Compile Proto files
Run the command below from the nats-streaming directory:

protoc -I pb/ pb/*.proto --go_out=plugins=grpc:pb

## Set up CockroachDB

#### Create user
cockroach user set shijuvar --insecure

#### Create Database
cockroach sql --insecure -e 'CREATE DATABASE ordersdb'

#### Grant privileges to the shijuvar user
cockroach sql --insecure -e 'GRANT ALL ON DATABASE ordersdb TO shijuvar'

### Start CockroachDB Cluster 

#### Start node 1:
cockroach start --insecure \
--store=ordersdb-1 \
--host=localhost \
--background

#### Start node 2:
cockroach start --insecure \
--store=ordersdb-2 \
--host=localhost \
--port=26258 \
--http-port=8081 \
--join=localhost:26257 \
--background

#### Start node 3:
cockroach start --insecure \
--store=ordersdb-3 \
--host=localhost \
--port=26259 \
--http-port=8082 \
--join=localhost:26257 \
--background

## Run NATS Streaming Server
nats-streaming-server \
--store file \
--dir ./data \
--max_msgs 0 \
--max_bytes 0

	

