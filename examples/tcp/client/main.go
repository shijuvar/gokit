package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// connect to this socket
	conn, err := net.Dial("tcp", "127.0.0.1:1433")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// read input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send messages to TCP server: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Reply from server: " + message)
	}
}
