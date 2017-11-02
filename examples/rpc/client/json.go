package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gorilla/rpc/json"

	"github.com/shijuvar/gokit/examples/rpc/rpcexample"
)

func main() {
	url := "http://localhost:1234/rpc"
	args := &rpcexample.Args{
		A: 2,
		B: 3,
	}
	message, err := json.EncodeClientRequest("Arith.Multiply", args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	var result rpcexample.Result
	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
