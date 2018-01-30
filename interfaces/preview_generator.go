package interfaces

type PreviewGenerator interface {
    GetPreviewLink(fileId string) (string, error)
}