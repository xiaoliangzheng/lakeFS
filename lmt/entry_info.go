package lmt

type EntryInfo interface {
	GetPath() string
	GetAddress() string
	GetSize() uint64
	GetChecksum() string
	GetMetadata() map[string]string
}
