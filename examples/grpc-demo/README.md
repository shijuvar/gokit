## Compile proto files
Run the command below from the grpc-demo directory:

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    customer/*.proto
