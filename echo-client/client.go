package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/smallinsky/grpc-echo/proto"
)

const (
	EchoServerAddr = "echo-server:8080"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("/cert/localhost.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err := grpc.Dial(EchoServerAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		r, err := c.Echo(ctx, &pb.EchoReq{Name: "ping"})
		if err != nil {
			log.Printf("could not echo: %v", err)
			time.Sleep(time.Second * 4)
			continue
		}
		log.Printf("echo response: %s", r.Name)
		time.Sleep(time.Second * 2)
	}
}
