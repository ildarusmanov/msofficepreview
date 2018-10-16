package services

import (
	"github.com/google/uuid"
	"github.com/ildarusmanov/msofficepreview/interfaces"
	"sync"
	"time"
)

type Token struct {
	Ttl       int64
	Timestamp int64
	Value     string
	FilePath  string
}

const (
	expireInterval = 100
)

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
	tokenLifetime int64
	tokens        *sync.Map
}

func CreateMemoryTokenProvider(tokenLifetime int64) *MemoryTokenProvider {
	p := &MemoryTokenProvider{
		tokenLifetime: tokenLifetime,
		tokens:        &sync.Map{},
	}

	go func() {
		for {
			time.Sleep(time.Duration(expireInterval) * time.Millisecond)
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
	if token, ok := p.tokens.Load(tokenValue); ok {
		return token.(*Token), true
	}

	return nil, false
}

func (p *MemoryTokenProvider) cleanUp() {
	p.tokens.Range(func(tokenValue, token interface{}) bool {
		if p.isExpired(token.(*Token)) {
			p.tokens.Delete(tokenValue)
		}

		return true
	})
}

func (p *MemoryTokenProvider) isExpired(token interfaces.Token) bool {
	return time.Now().Unix() > token.GetTtl()
}

func (p *MemoryTokenProvider) tokenExists(tokenValue string) bool {
	_, ok := p.tokens.Load(tokenValue)

	return ok
}

func (p *MemoryTokenProvider) getTokenTtl() int64 {
	return time.Now().Unix() + p.tokenLifetime
}

func (p *MemoryTokenProvider) createNewToken(filePath string) interfaces.Token {
	newToken := CreateToken(filePath, p.getTokenTtl())

	p.tokens.Store(newToken.Value, newToken)

	return newToken
}
