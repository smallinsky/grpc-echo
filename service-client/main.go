package main

import (
	"log"
	"time"

	pb "github.com/smallinsky/service-mesh/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Echo(ctx, &pb.EchoReq{Name: "ping"})
		if err != nil {
			log.Fatalf("could not echo: %v", err)
		}
		log.Printf("echo response: %s", r.Name)
		time.Sleep(time.Second * 2)
	}
}
