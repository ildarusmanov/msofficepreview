package interfaces

type PreviewGenerator interface {
	GetPreviewLink(filePath string) (PreviewInfo, error)
}
