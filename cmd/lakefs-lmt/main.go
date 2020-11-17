package main

import (
	"flag"
	"github.com/treeverse/lakefs/lmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("l", ":8000", "listen address")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	lmt.RegisterLakeMetadataTrackerServer(s, &lmt.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
