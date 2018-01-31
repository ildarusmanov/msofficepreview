package interfaces

type TokenProvider interface {
	Generate() string
	Validate(accessToken string) bool
}
