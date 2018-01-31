package interfaces

type FileInfo interface {
	GetFileName() string
	GetSize() int64
	GetVersion() string
	GetOwnerId() string
}
