package interfaces

type TokenProvider interface {
	Generate(filePath string) string
	Validate(token Token) bool
	FindToken(accessToken string) (Token, bool)
}
