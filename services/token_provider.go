package services

type TokenProvider struct{}

func CreateTokenProvider() *TokenProvider {
    return &TokenProvider{}
}

func Generate() string {
    return "new-token"
}

func Validate(token) bool {
    return "new-token" == token
}
