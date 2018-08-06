package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	pb "github.com/smallinsky/grpc-echo/proto"
)

type server struct{}

func (s *server) Echo(ctx context.Context, req *pb.EchoReq) (*pb.EchoResp, error) {
	log.Printf("get echo request %s", req.GetName())
	return &pb.EchoResp{
		Name: req.GetName(),
	}, nil

}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("failed to load TLS certs")
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	pb.RegisterEchoServer(s, &server{})

	reflection.Register(s)
	log.Print("echo service start")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
