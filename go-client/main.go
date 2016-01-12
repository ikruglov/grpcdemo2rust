package main

import (
	"flag"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "./helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	str := flag.String("str", defaultName, "your string")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: *str})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
