package design

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Protocol string
	Addr     string
	Port     int
	Timeout  time.Duration
	MaxConn  int
	TLS      *tls.Config
}

type Option func(*Server)

func Timeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

func TLS(tls *tls.Config) Option {
	return func(server *Server) {
		server.TLS = tls
	}
}

func Protocol(protocol string) Option {
	return func(server *Server) {
		server.Protocol = protocol
	}
}

func MaxConn(conn int) Option {
	return func(server *Server) {
		server.MaxConn = conn
	}
}

func Port(port int) Option {
	return func(server *Server) {
		server.Port = port
	}
}

// build模式
// config模式
// Function optional模式

func NewServer(addr string, port int, options ...func(*Server)) *Server {
	server := &Server{
		Protocol: "tcp",
		Addr:     addr,
		Port:     port,
		Timeout:  30 * time.Second,
		MaxConn:  1000,
		TLS:      nil,
	}

	for _, option := range options {
		option(server)
	}
	return server
}
