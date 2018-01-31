package services

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type Token struct {
	Timestamp int64
	Value     string
}

func CreateToken() *Token {
	tokenValue := uuid.New().URN()
	timestamp := time.Now().Unix()

	return &Token{
		Timestamp: timestamp,
		Value:     tokenValue,
	}
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

func (p *MemoryTokenProvider) Generate() string {
	newToken := p.createNewToken()

	return newToken.Value
}

func (p *MemoryTokenProvider) Validate(tokenValue string) bool {
	return p.tokenExists(tokenValue)
}

func (p *MemoryTokenProvider) cleanUp() {
	for tokenValue, token := range p.tokens {
		if p.isExpired(token.Timestamp) {
			delete(p.tokens, tokenValue)
		}
	}
}

func (p *MemoryTokenProvider) isExpired(tokenTime int64) bool {
	return time.Now().Unix() > tokenTime+p.tokenLifetime
}

func (p *MemoryTokenProvider) tokenExists(tokenValue string) bool {
	p.Lock()
	_, ok := p.tokens[tokenValue]
	p.Unlock()

	return ok
}

func (p *MemoryTokenProvider) createNewToken() *Token {
	newToken := CreateToken()

	p.Lock()
	p.tokens[newToken.Value] = newToken
	p.Unlock()

	return newToken
}
