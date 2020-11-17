package lmt

import (
	"context"
	"fmt"
)

type Server struct {
	LakeMetadataTrackerServer
}

func (s *Server) GetEntry(ctx context.Context, in *GetEntryRequest) (*GetEntryResponse, error) {
	fmt.Println("Get entry was called")
	return &GetEntryResponse{}, nil
}
