package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/treeverse/lakefs/lmt"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("s", "localhost:8000", "server address")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLakeMetadataTrackerClient(conn)

	ctx := context.Background()
	r, err := c.GetEntry(ctx, &pb.GetEntryRequest{
		Repository: "repo1",
		Ref:        "master",
		Path:       "path1",
	})
	if err != nil {
		log.Fatalf("could not get entry: %s", err)
	}
	log.Printf("Done!!!!: %v", r)
}
