package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "./helloworld"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("new request", in.Name)

	buf, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	res, err := http.Post("http://localhost:3000/", "application/x-protobuf", bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var reply pb.HelloReply
	if err = proto.Unmarshal(result, &reply); err != nil {
		return nil, err
	}

	log.Println("response", reply.Message)
	return &reply, nil
}

func main() {
	log.Println("start grpc service")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
