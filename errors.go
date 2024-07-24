package go_wahoo

import "errors"

var (
	ErrInvalidRedirectURI     = errors.New("invalid redirect uri")
	ErrInvalidScopes          = errors.New("invalid scopes")
	ErrInvalidClientID        = errors.New("client id is required")
	ErrInvalidClientSecret    = errors.New("client secret is required")
	ErrFailedToGetAccessToken = errors.New("failed to get access token")
)
