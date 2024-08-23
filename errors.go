package go_wahoo

import "errors"

type RequestError struct {
	Err   error
	Code  int
	Debug string
}

func NewError(err error, code int, debug string) *RequestError {
	return &RequestError{
		Err:   err,
		Code:  code,
		Debug: debug,
	}
}

var (
	ErrInvalidRedirectURI     = errors.New("invalid redirect uri")
	ErrInvalidScopes          = errors.New("invalid scopes")
	ErrInvalidClientID        = errors.New("client id is required")
	ErrInvalidClientSecret    = errors.New("client secret is required")
	ErrFailedToGetAccessToken = errors.New("failed to get access token")
	ErrInvalidRefreshToken    = errors.New("invalid refresh token")
	ErrInvalidCode            = errors.New("invalid code")
)
