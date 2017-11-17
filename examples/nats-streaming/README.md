## Technologies Used: 
* NATS Streaming
* gRPC
* CockroachDB

## Components in the Demo App
* pb: Protocol Buffers definitions to describe message types and RPC endpoints.
orderservice: An HTTP API server that let customers to create Orders. When a new Order is placed, an event “OrderCreated” is triggered, hence it calls an gRPC method “CreateEvent” provided by eventstore to publish events to the Event Store.
* eventstore: A gRPC server and a NATS Streaming client that persists domain events into Event Store and publish events on NATS Streaming channels. This example assumes that state of the application is composed by various events ( A fluid implementation of Event Sourcing pattern). All command operations are persisted into an Event Store as events. Here CockroachDB is used for persisting events.
* restuarantservice: A NATS Streaming client that subscribe messages from a NATS Streaming channel “order-notification” to get messages when new orders are created via orderservice and messages are published over channel “order-notification” from eventstore.
* orderquery-store1: A NATS Streaming client that subscribes messages with a QueueGroup (a NATS messaging pattern) from a NATS Streaming channel “order-notification” to get messages when events are happened on a aggregate Order. The objective of this package is to persist data model for querying data, based on the domain events persisted in the Event Store. The example demo assumes that separate data models are being used for both command operations and query operations (CQRS). Because you’re keeping separate data models for both command and query, you can have denormalized data sets o n the data models for query. Here CockroachDB is used for persisting data sets for query model. In real-world scenarios, separate databases will be used for both command and query models.
* orderquery-store2: A NATS Streaming client that subscribes messages with a QueueGroup from a NATS Streaming channel “order-notification”. Both orderquery-store1 and orderquery-store2 do the same thing — perform the data replication logic for making a store for querying the data which is constructed from Event Store. In order to distribute data replication logic, it works as QueueGroup subscriber clients (orderquery-store1 and orderquery-store2).
* store: This is a shared library package that provides persistence logic to working with CockroachDB database. 

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

	

