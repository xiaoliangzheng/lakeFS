package lmt

import context "context"

type Committed interface {
	GetEntry(ctx context.Context, repository string, branch string, path string) (error, EntryInfo)
}
