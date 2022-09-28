package constant

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
)

// TokenInfo interface
type TokenInfo interface {
	GetAccessToken() string
	GetRefreshToken() string
	GetTokenType() string
	EncodeToJSON() ([]byte, error)
}

type tokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}
type Token struct {
	AccessToken  string    `json:"access_token,omitempty"`
	UserID       uuid.UUID `json:"user_id,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
}

// GetAccessToken return access token
func (t *tokenInfo) GetAccessToken() string {
	return t.AccessToken
}

// GetRefreshToken return refresh token
func (t *tokenInfo) GetRefreshToken() string {
	return t.RefreshToken
}

// GetTokenType return token type
func (t *tokenInfo) GetTokenType() string {
	return t.TokenType
}

func (t *tokenInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}
