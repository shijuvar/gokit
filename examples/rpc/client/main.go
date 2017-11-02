package main

import (
	"log"
	"net/rpc"

	"github.com/shijuvar/gokit/examples/rpc/rpcexample"
)

func main() {
	//make connection to rpc server
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &rpcexample.Args{
		A: 2,
		B: 3,
	}
	//this will store returned result
	var result rpcexample.Result
	//call remote procedure with args
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//we got our result in result
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
