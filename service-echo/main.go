package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/smallinsky/service-mesh/proto"
)

type server struct{}

func (s *server) Echo(ctx context.Context, req *pb.EchoReq) (*pb.EchoResp, error) {
	log.Printf("get EchoReq{Name: %s}", req.GetName())
	return &pb.EchoResp{
		Name: req.GetName(),
	}, nil

}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &server{})

	log.Print("echo service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
