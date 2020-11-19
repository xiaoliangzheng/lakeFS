package lmt

import (
	"context"
)

type Service struct {
	LakeMetadataTrackerServer

	Committed   Committed
	Uncommitted Uncommitted
}

func (s *Service) PutEntry(ctx context.Context, r *PutEntryRequest) (*PutEntryResponse, error) {
	err = s.Uncommitted.Repository(r.Repository).Branch(r.Branch).PutEntry(ctx, &r.Entry)
	if err != nil {
		return nil, err
	}
	return &PutEntryResponse{}
}

func (s *Service) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) Log(context.Context, *LogRequest) (*LogResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) GetBranch(context.Context, *GetBranchRequest) (*GetBranchResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) ListBranches(context.Context, *ListBranchesRequest) (*ListBranchesResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) Merge(context.Context, *MergeRequest) (*MergeResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	panic("TODO(barak): implement")
}

func (s *Service) Revert(context.Context, *RevertRequest) (*RevertResponse, error) {
	panic("TODO(barak): implement")
}
