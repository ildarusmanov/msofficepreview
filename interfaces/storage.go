package interfaces

type Storage interface {
	GetFileInfo(fileId string) (FileInfo, error)
	GetContents(fileId string) ([]byte, error)
}
