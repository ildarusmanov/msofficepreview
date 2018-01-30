package services

import (
    "github.com/ildarusmanov/msofficepreview/interfaces"
)

type PreviewGenerator struct {
    storage interfaces.Storage
    tokenProvider interfaces.TokenProvider
}

func CreatePreviewGenerator(tokenProvider interfaces.TokenProvider, storage interfaces.Storage) *PreviewGenerator {
    return &PreviewGenerator{
        tokenProvider: tokenProvider,
        storage: storage,
    }
}

func (g *PreviewGenerator) GetPreviewLink(fileId string) (string, error) {
    _, err := g.getFileInfo(fileId)

    if err != nil {
        return "", err
    }

    return g.buildPreviewLink(g.generateToken(), fileId), nil
}

func (g *PreviewGenerator) getFileInfo(fileId string) (interfaces.FileInfo, error) {
    return g.storage.GetFileInfo(fileId)
}

func (g *PreviewGenerator) generateToken() string {
    return g.tokenProvider.Generate()
}

func (g *PreviewGenerator) buildPreviewLink(accessToken, fileId string) string {
    return ""
}
