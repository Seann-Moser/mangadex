package auth

import (
	"context"
	"errors"
	"sync"
	"time"
)

type memoryEntry struct {
	tokens    OAuthTokens
	creds     *OAuthCredentials
	expiresAt time.Time
}

var _ TokenStore = (*InMemoryTokenStore)(nil)

type InMemoryTokenStore struct {
	mu    sync.RWMutex
	store map[string]memoryEntry
}

func NewInMemoryTokenStore() *InMemoryTokenStore {
	return &InMemoryTokenStore{
		store: make(map[string]memoryEntry),
	}
}

func (m *InMemoryTokenStore) Save(
	_ context.Context,
	userKey string,
	tokens OAuthTokens,
	creds *OAuthCredentials,
	expiresAt time.Time,
) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	// defensive copy of creds
	var storedCreds *OAuthCredentials
	if creds != nil {
		c := *creds
		storedCreds = &c
	}

	m.store[userKey] = memoryEntry{
		tokens:    tokens,
		creds:     storedCreds,
		expiresAt: expiresAt,
	}

	return nil
}

func (m *InMemoryTokenStore) Load(
	_ context.Context,
	userKey string,
) (*OAuthTokens, *OAuthCredentials, time.Time, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	entry, ok := m.store[userKey]
	if !ok {
		return nil, nil, time.Time{}, errors.New("token not found")
	}

	if time.Now().After(entry.expiresAt) {
		delete(m.store, userKey)
		return nil, nil, time.Time{}, errors.New("token expired")
	}

	tokens := entry.tokens

	var creds *OAuthCredentials
	if entry.creds != nil {
		c := *entry.creds
		creds = &c
	}

	return &tokens, creds, entry.expiresAt, nil
}
