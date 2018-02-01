package services

import (
    "github.com/ildarusmanov/msofficepreview/interfaces"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Token struct {
    Ttl       int64
	Timestamp int64
	Value     string
    FilePath  string
}

func CreateToken(filePath string, ttl int64) *Token {
	tokenValue := uuid.New().URN()
	timestamp := time.Now().Unix()

	return &Token{
        Ttl:       ttl,
		Timestamp: timestamp,
		Value:     tokenValue,
        FilePath:  filePath,
	}
}

func (t *Token) GetValue() string {
    return t.Value
}

func (t *Token) GetFilePath() string {
    return t.FilePath
}

func (t *Token) GetTtl() int64 {
    return t.Ttl
}

type MemoryTokenProvider struct {
	sync.Mutex
	tokenLifetime int64
	tokens        map[string]*Token
}

func CreateMemoryTokenProvider(tokenLifetime int64) *MemoryTokenProvider {
	p := &MemoryTokenProvider{
		tokenLifetime: tokenLifetime,
		tokens:        make(map[string]*Token),
	}

	go func() {
		for {
			time.Sleep(time.Duration(100) * time.Millisecond)
			p.cleanUp()
		}
	}()

	return p
}

func (p *MemoryTokenProvider) Generate(filePath string) string {
	newToken := p.createNewToken(filePath)

	return newToken.GetValue()
}

func (p *MemoryTokenProvider) Validate(token interfaces.Token) bool {
	return !p.isExpired(token)
}

func (p *MemoryTokenProvider) FindToken(tokenValue string) (interfaces.Token, bool) {
    p.Lock()
    token, ok := p.tokens[tokenValue]
    p.Unlock()

    return token, ok
}

func (p *MemoryTokenProvider) cleanUp() {
	for tokenValue, token := range p.tokens {
		if p.isExpired(token) {
			delete(p.tokens, tokenValue)
		}
	}
}

func (p *MemoryTokenProvider) isExpired(token interfaces.Token) bool {
	return time.Now().Unix() > token.GetTtl()
}

func (p *MemoryTokenProvider) tokenExists(tokenValue string) bool {
	p.Lock()
	_, ok := p.tokens[tokenValue]
	p.Unlock()

	return ok
}

func (p *MemoryTokenProvider) getTokenTtl() int64 {
    return time.Now().Unix() + p.tokenLifetime
}

func (p *MemoryTokenProvider) createNewToken(filePath string) interfaces.Token {
	newToken := CreateToken(filePath, p.getTokenTtl())

	p.Lock()
	p.tokens[newToken.Value] = newToken
	p.Unlock()

	return newToken
}
