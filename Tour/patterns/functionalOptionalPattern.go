package main

import "fmt"

type Server struct {
	port    int
	timeout int
	// ... other config
}

type ServerOption func(*Server)

func WithPort(port int) ServerOption {
	return func(s *Server) { // This is a closure over 'port'
		s.port = port
	}
}

func WithTimeout(timeout int) ServerOption {
	return func(s *Server) { // This is a closure over 'timeout'
		s.timeout = timeout
	}
}

func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		port:    8080,
		timeout: 30, // default
	}
	for _, opt := range opts {
		opt(s) // Each opt is a function that modifies 's'
	}
	return s
}

func main() {
	// Create a server with custom port and timeout using functional options
	server := NewServer(
		WithPort(9000),
		WithTimeout(60),
	)
	fmt.Printf("Server config: %+v\n", server)
}
