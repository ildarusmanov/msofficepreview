package interfaces

type PreviewGenerator interface {
	GetPreviewLink(fileId string) (PreviewInfo, error)
}
