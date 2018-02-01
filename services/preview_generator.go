package services

import (
	"github.com/ildarusmanov/msofficepreview/interfaces"
	"net/url"
	"path"
	"strings"
)

type PreviewInfo struct {
	src      string
	token    string
	tokenTtl int64
}

func CreatePreviewInfo(src, token string, tokenTtl int64) *PreviewInfo {
	return &PreviewInfo{
		src:      src,
		token:    token,
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
	serverHost    string
	wopiDiscovery interfaces.WopiDiscovery
	storage       interfaces.Storage
	tokenProvider interfaces.TokenProvider
}

func CreatePreviewGenerator(
	serverHost string,
	wopiDiscovery interfaces.WopiDiscovery,
	tokenProvider interfaces.TokenProvider,
	storage interfaces.Storage,
) *PreviewGenerator {
	return &PreviewGenerator{
		serverHost:    serverHost,
		wopiDiscovery: wopiDiscovery,
		tokenProvider: tokenProvider,
		storage:       storage,
	}
}

func (g *PreviewGenerator) GetPreviewLink(filePath string) (interfaces.PreviewInfo, error) {
	_, err := g.getFileInfo(filePath)

	if err != nil {
		return nil, err
	}

	return g.buildPreviewInfo(g.generateToken(filePath), filePath)
}

func (g *PreviewGenerator) getFileInfo(filePath string) (interfaces.FileInfo, error) {
	return g.storage.GetFileInfo(filePath)
}

func (g *PreviewGenerator) generateToken(filePath string) interfaces.Token {
	tokenValue := g.tokenProvider.Generate(filePath)
	token, _ := g.tokenProvider.FindToken(tokenValue)

	return token
}

func (g *PreviewGenerator) getActionUrlsrc(ext string) (string, error) {
	return g.wopiDiscovery.FindPreviewUrl("internal-https", ext)
}

func (g *PreviewGenerator) buildPreviewInfo(token interfaces.Token, filePath string) (*PreviewInfo, error) {
	ext := strings.TrimPrefix(path.Ext(filePath), ".")
	urlsrc, err := g.getActionUrlsrc(ext)

	if err != nil {
		return nil, err
	}

	WOPIsrc := url.QueryEscape(g.serverHost + "/wopi/files/" + token.GetValue())
	urlsrc = urlsrc + "&WOPIsrc=" + WOPIsrc

	return CreatePreviewInfo(urlsrc, token.GetValue(), token.GetTtl()), nil
}
