package go_wahoo

type IWahoo interface {
	GetAuthenticateURL() (*string, error)
	GetAccessToken(code string) (*TokenResponse, *RequestError)
	RefreshToken(refreshToken string) (*TokenResponse, *RequestError)
}
