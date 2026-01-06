package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	tokenURL  = "https://auth.mangadex.org/realms/mangadex/protocol/openid-connect/token"
	apiDomain = "api.mangadex.org"
)

type OAuthClient struct {
	HTTPClient *http.Client

	// defaults (used when user has none)
	DefaultCreds OAuthCredentials

	UserKey string
	Store   TokenStore

	mu        sync.Mutex
	tokens    OAuthTokens
	creds     OAuthCredentials
	expiresAt time.Time
}

func (c *OAuthClient) Login(ctx context.Context, username string, password string, creds *OAuthCredentials) error {
	if creds != nil && creds.UserKey == "" {
		return fmt.Errorf("missing user key")
	}
	form := url.Values{}
	form.Set("grant_type", "password")
	form.Set("username", username)
	form.Set("password", password)
	form.Set("client_id", creds.ClientID)
	form.Set("client_secret", creds.ClientSecret)
	return c.doTokenRequest(ctx, form, creds)
}

func (c *OAuthClient) ApplyAuth(userKey string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		if !strings.Contains(req.URL.Host, apiDomain) {
			return nil
		}

		tok, cred, exp, err := c.LoadFromStore(ctx, userKey)
		if err != nil {
			return err
		}
		expired := time.Now().After(exp)

		if expired {
			if err := c.refresh(ctx, tok, cred); err != nil {
				return err
			}
		}

		if tok.AccessToken == "" {
			return errors.New("no access token")
		}

		req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
		return nil
	}
}

func (c *OAuthClient) storeTokens(tok OAuthTokens, creds *OAuthCredentials) {
	if c.Store != nil {
		_ = c.Store.Save(
			context.Background(),
			creds.UserKey,
			tok,
			creds,
			time.Now().Add(time.Duration(tok.ExpiresIn-30)*time.Second),
		)
	} else {

	}
}

func (c *OAuthClient) LoadFromStore(ctx context.Context, key string) (*OAuthTokens, *OAuthCredentials, time.Time, error) {
	if c.Store == nil {
		return nil, nil, time.Now(), nil
	}

	tok, userCreds, exp, err := c.Store.Load(ctx, key)
	if err != nil {
		return nil, nil, time.Now(), err
	}

	return tok, c.resolveCreds(userCreds), exp, nil
}

func (c *OAuthClient) resolveCreds(userCreds *OAuthCredentials) *OAuthCredentials {
	if userCreds != nil &&
		userCreds.ClientID != "" &&
		userCreds.ClientSecret != "" {
		return userCreds
	}
	return &c.DefaultCreds
}

func (c *OAuthClient) refresh(ctx context.Context, tok *OAuthTokens, creds *OAuthCredentials) error {
	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", tok.RefreshToken)
	form.Set("client_id", creds.ClientID)
	form.Set("client_secret", creds.ClientSecret)

	return c.doTokenRequest(ctx, form, creds)
}

func (c *OAuthClient) doTokenRequest(ctx context.Context, form url.Values, creds *OAuthCredentials) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		tokenURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("oauth request failed")
	}

	var tok OAuthTokens
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return err
	}

	c.storeTokens(tok, creds)
	return nil
}
