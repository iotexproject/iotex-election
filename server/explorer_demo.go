// This file shoule be replaced with proper test case file and
// which will take care of starting the server through proper route.

// For now, this code is NOT working, after code restructuring.

package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ashishsnigam/iotex-election/explorer_pb"
)

func RunServer() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterExplorerServer(s, &pb.server{})
	s.Serve(lis)
}
