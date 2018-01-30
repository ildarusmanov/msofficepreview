package services

import (
    "time"
    "sync"
    "github.com/google/uuid"
)

type Token struct {
    Timestamp int64
    Value string
}

func CreateToken() *Token {
    tokenValue := uuid.New().URN()
    timestamp := time.Now().Unix()

    return &Token{
        Timestamp: timestamp,
        Value: tokenValue,
    }
}

type TokenProvider struct {
    sync.Mutex
    tokenLifetime int64
    tokens map[string]*Token
}

func CreateTokenProvider(tokenLifetime int64) *TokenProvider {
    return &TokenProvider{
        tokenLifetime: tokenLifetime,
        tokens: make(map[string]*Token),
    }
}

func (p *TokenProvider) Generate() string {
    newToken := p.createNewToken()

    return newToken.Value
}

func (p* TokenProvider) Validate(tokenValue string) bool {
    return p.tokenExists(tokenValue)
}

func (p *TokenProvider) CleanUp() {
    for tokenValue, token := range p.tokens {
        if p.isExpired(token.Timestamp) {
            delete(p.tokens, tokenValue)
        }
    }
}

func (p *TokenProvider) isExpired(tokenTime int64) bool {
    return time.Now().Unix() > tokenTime + p.tokenLifetime
}

func (p *TokenProvider) tokenExists(tokenValue string) bool {
    p.Lock()
    _, ok := p.tokens[tokenValue]
    p.Unlock()

    return ok
}

func (p *TokenProvider) createNewToken() *Token {
    newToken := CreateToken()

    p.Lock()
    p.tokens[newToken.Value] = newToken
    p.Unlock()

    return newToken
}
