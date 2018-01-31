package services

import (
    "github.com/ildarusmanov/msofficepreview/interfaces"
)

type PreviewInfo struct {
    src string
    token string
    tokenTtl int64
}

func CreatePreviewInfo(src, token string, tokenTtl int64) *PreviewInfo {
    return &PreviewInfo{
        src: src,
        token: token,
        tokenTtl: tokenTtl,
    }
}

func (p *PreviewInfo) GetSrc() string {
    return p.src
}

func (p *PreviewInfo) GetToken() string {
    return p.token
}

func (p *PreviewInfo) GetTokenTtl() int64 {
    return p.tokenTtl
}

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

func (g *PreviewGenerator) GetPreviewLink(fileId string) (*PreviewInfo, error) {
    _, err := g.getFileInfo(fileId)

    if err != nil {
        return nil, err
    }

    return g.buildPreviewInfo(g.generateToken(), fileId), nil
}

func (g *PreviewGenerator) getFileInfo(fileId string) (interfaces.FileInfo, error) {
    return g.storage.GetFileInfo(fileId)
}

func (g *PreviewGenerator) generateToken() string {
    return g.tokenProvider.Generate()
}

func (g *PreviewGenerator) buildPreviewInfo(accessToken, fileId string) *PreviewInfo {
    return CreatePreviewInfo("/wopi/files"+ fileId, accessToken, int64(0))
}
