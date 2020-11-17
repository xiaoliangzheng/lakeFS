package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/treeverse/lakefs/lmt"
	"google.golang.org/grpc"
)

type server struct {
	pb.LakeMetadataTrackerServer
}

func (s *server) GetEntry(ctx context.Context, in *pb.GetEntryRequest) (*pb.GetEntryReply, error) {
	fmt.Println("Get entry was called")
	return &pb.GetEntryReply{}, nil
}

func main() {
	addr := flag.String("l", ":8000", "listen address")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLakeMetadataTrackerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
