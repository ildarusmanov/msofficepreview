package services

// preview generator
type PreviewGenerator struct {
  accessToken string
}

// create new preview link generator
// with given accessToken string
func CreatePreviewGenerator(accessToken string) *PreviewGenerator {
  return &PreviewGenerator{accessToken}
}

// get current access token
func (pg *PreviewGenerator) GetAccessToken() string {
  return pg.accessToken
}

// generate preview link by given link
func (pg *PreviewGenerator) GetPreviewLink(link string) string {
  return "http://fake.link.test"
}

