package main

import (
	"flag"
	"log"
	"net"

	"github.com/treeverse/lakefs/lmt"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("l", ":8000", "listen address")
	flag.Parse()

	log.Printf("lakeFS lmt server - listen on %s", *addr)
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	lmtService := &lmt.Service{}
	s := grpc.NewServer()
	lmt.RegisterLakeMetadataTrackerServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
