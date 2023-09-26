package main

import "fmt"

type Server struct {
	host     string
	port     int
	protocol string
}

type ServerOption func(server *Server)

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}
func WithProtocol(protocol string) ServerOption {
	return func(s *Server) {
		s.protocol = protocol
	}
}

func NewServer(host string, opts ...ServerOption) *Server {
	server := &Server{
		host:     host,
		port:     8080,
		protocol: "http",
	}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

func main() {
	server1 := NewServer("localhost") // with default options
	fmt.Printf("sever 1:%+v\n", *server1)
	server2 := NewServer("localhost", WithPort(3000), WithProtocol("https"))
	fmt.Printf("sever 2:%+v\n", *server2)
}
