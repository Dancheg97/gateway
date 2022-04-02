package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"gateway/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out pb/ api.proto

type server struct {
	pb.UnimplementedGatewayServer
}

func (s *server) PostExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Println(in)
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) GetExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Println(in)
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) DeleteExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Println(in)
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) PutExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Println(in)
	return &pb.Message{Id: in.Id}, nil
}

func (s *server) PatchExample(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Println(in)
	return &pb.Message{Id: in.Id}, nil
}

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterGatewayHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func runGrpc() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	go runGrpc()
	runRest()
}
