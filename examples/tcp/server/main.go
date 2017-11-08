package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings" // only needed below for sample processing
)

func main() {

	fmt.Println("Launching TCP server...")

	// listen
	ln, err := net.Listen("tcp", ":1433")
	if err != nil {
		log.Fatal(err)
	}

	// accept connection on port
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// processing the message string
		newMessage := strings.ToTitle(message)
		// send new string back to client
		conn.Write([]byte(newMessage + "\n"))
	}
}
