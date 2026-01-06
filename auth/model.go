package auth

type OAuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

type OAuthCredentials struct {
	UserKey      string `json:"user_key"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
