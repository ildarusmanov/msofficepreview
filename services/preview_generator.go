package services

import (
    "strings"
    "path"
    "net/url"
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
    serverHost string
    wopiDiscovery interfaces.WopiDiscovery
    storage interfaces.Storage
    tokenProvider interfaces.TokenProvider
}

func CreatePreviewGenerator(
    serverHost string,
    wopiDiscovery interfaces.WopiDiscovery,
    tokenProvider interfaces.TokenProvider,
    storage interfaces.Storage,
) *PreviewGenerator {
    return &PreviewGenerator{
        serverHost: serverHost,
        wopiDiscovery: wopiDiscovery,
        tokenProvider: tokenProvider,
        storage: storage,
    }
}

func (g *PreviewGenerator) GetPreviewLink(fileId string) (*PreviewInfo, error) {
    _, err := g.getFileInfo(fileId)

    if err != nil {
        return nil, err
    }

    return g.buildPreviewInfo(g.generateToken(), fileId)
}

func (g *PreviewGenerator) getFileInfo(fileId string) (interfaces.FileInfo, error) {
    return g.storage.GetFileInfo(fileId)
}

func (g *PreviewGenerator) generateToken() string {
    return g.tokenProvider.Generate()
}

func (g *PreviewGenerator) getActionUrlsrc(ext string) (string, error) {
    return g.wopiDiscovery.FindPreviewUrl("internal-https", ext)
}

func (g *PreviewGenerator) buildPreviewInfo(accessToken, fileId string) (*PreviewInfo, error) {
    ext := strings.TrimPrefix(path.Ext(fileId), ".")
    urlsrc, err := g.getActionUrlsrc(ext)

    if err != nil {
        return nil, err
    }

    WOPIsrc := url.QueryEscape(g.serverHost + "/wopi/files/" + fileId)
    urlsrc = urlsrc + "?WOPIsrc="+ WOPIsrc

    return CreatePreviewInfo(urlsrc, accessToken, int64(0)), nil
}
