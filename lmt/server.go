package lmt

import (
	"context"
)

type Server struct {
	LakeMetadataTrackerServer

	Committed   Committed
	Uncommitted Uncommitted
}

func (s *Server) PutEntry(context.Context, *PutEntryRequest) (*PutEntryResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) Log(context.Context, *LogRequest) (*LogResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) GetBranch(context.Context, *GetBranchRequest) (*GetBranchResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) ListBranches(context.Context, *ListBranchesRequest) (*ListBranchesResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) Merge(context.Context, *MergeRequest) (*MergeResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Server) Revert(context.Context, *RevertRequest) (*RevertResponse, error) {
	panic("TODO(barak): implement")
}
