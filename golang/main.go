package main

import (
	"context"
	"log"
	"net"

	"github.com/atlanhq/grpc-example/golang/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3000"
)

// define a server struct
type server struct{}

// implement the interface
func (s *server) SayHello(ctx context.Context, in *pb.Hi) (*pb.Bye, error) {
	log.Printf("Received: %+v", in.Msg)
	return &pb.Bye{Msg: "Bye Bye", Id: "dsiabdxasbdbasjbdkjbsadb"}, nil
}

func main() {
	// listen on tcp
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	// attach a grpc server to listener
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	// enable server reflection
	reflection.Register(s)

	log.Println("listening...")
	log.Fatal(s.Serve(lis))
}
