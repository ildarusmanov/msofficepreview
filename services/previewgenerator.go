package services

type PreviewGenerator struct {
  accessToken string
}

func CreatePreviewGenerator(accessToken string) *PreviewGenerator {
  return &PreviewGenerator{accessToken}
}

func (pg *PreviewGenerator) GetAccessToken() string {
  return pg.accessToken
}

func (pg *PreviewGenerator) GetPreviewLink(link string) string {
  return "http://fake.link.test"
}

