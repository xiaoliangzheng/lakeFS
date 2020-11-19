package lmt

import context "context"

type UncommittedBranch interface {
	PutEntry(ctx context.Context, entry EntryInfo)
}

type UncommittedRepository interface {
	Branch(branch string) UncommittedBranch
}

type Uncommitted interface {
	Repository(name string) UncommittedBranch
	//PutEntry(ctx context.Context, repository string, branch string, entry EntryInfo) error
}
