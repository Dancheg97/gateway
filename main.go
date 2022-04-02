package main

import (
	"context"
	"log"
	"net"

	"gateway/pb"

	"google.golang.org/grpc"
)

// protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out pb/ api.proto

type server struct {
	pb.UnimplementedGatewayServer
}

func (s *server) PostExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) GetExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) DeleteExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) PutExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{Id: in.Id}, nil
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
