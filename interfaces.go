package go_wahoo

type IWahoo interface {
	GetAuthenticateURL(uniqueCode string) (*string, error)
	GetAccessToken(code, uniqueCode string) (*TokenResponse, *RequestError)
	RefreshToken(refreshToken, uniqueCode string) (*TokenResponse, *RequestError)
	GetAllWorkout(token string, page int, limit int) (*WorkoutsResponse, *RequestError)
	DeAuthorize(token string) *RequestError
}
