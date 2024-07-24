package go_wahoo

type IWahoo interface {
	GetAuthenticateURL() (*string, error)
	GetAccessToken(code string) (*string, error)
}
