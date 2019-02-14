package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/ashishsnigam/iotex-election/explorer_pb"
)

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterExplorerServer(s, &pb.server{})
	s.Serve(lis)
}
