package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/shijuvar/gokit/examples/rpc/rpcexample"
)

func main() {
	//register Arith object as a service
	arith := new(rpcexample.Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatalf("Format of service Arith isn't correct. %s", err)
	}
	rpc.HandleHTTP()
	//start listening for messages on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
