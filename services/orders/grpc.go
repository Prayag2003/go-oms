package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	address string
}

func NewGRPCServer(address string) *GRPCServer {
	return &GRPCServer{
		address: address,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	grpcServer := grpc.NewServer()

	// register the grpc services here

	log.Println("Starting GRPC Server on ", s.address, " ...")
	return grpcServer.Serve(lis)
}
