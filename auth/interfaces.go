package auth

import (
	"context"
	"time"
)

type TokenStore interface {
	Load(ctx context.Context, userKey string) (
		tokens *OAuthTokens,
		creds *OAuthCredentials,
		expiresAt time.Time,
		err error,
	)

	Save(
		ctx context.Context,
		userKey string,
		tokens OAuthTokens,
		creds *OAuthCredentials,
		expiresAt time.Time,
	) error
}
