package go_wahoo

import (
	"encoding/json"
	"time"
)

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"createdAt"`
}

type ITokenResponse interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() time.Time
	GetRefreshToken() string
	GetScope() string
	GetCreatedAt() time.Time
}

func UnmarshalToResponse(data []byte) (*TokenResponse, *RequestError) {
	var resp TokenResponse
	err := json.Unmarshal(data, &resp)

	if err != nil {
		return nil, NewError(err, 500, "failed to unmarshal response")
	}

	return &resp, nil
}

func (t TokenResponse) GetAccessToken() string {
	return t.AccessToken
}

func (t TokenResponse) GetTokenType() string {
	return t.TokenType
}

func (t TokenResponse) GetExpiresAt() time.Time {
	if t.ExpiresIn == 0 {
		return time.Time{}
	}

	return time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)
}

func (t TokenResponse) GetRefreshToken() string {
	return t.RefreshToken
}

func (t TokenResponse) GetScope() string {
	return t.Scope
}

func (t TokenResponse) GetCreatedAt() time.Time {
	return time.Unix(t.CreatedAt, 0)
}
