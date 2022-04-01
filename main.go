package main

import (
	"context"
	"log"
	"net"

	"gateway/pb"

	"google.golang.org/grpc"
)

//protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. api.proto

type server struct {
	pb.UnimplementedGatewayServer
}

func (s *server) Echo(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{Message: "hey friend"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
