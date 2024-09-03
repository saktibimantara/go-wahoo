package go_wahoo

type IWahoo interface {
	GetAuthenticateURL(uniqueCode string) (*string, error)
	GetAccessToken(code, uniqueCode string) (*TokenResponse, *RequestError)
	RefreshToken(refreshToken, uniqueCode string) (*TokenResponse, *RateLimit, *RequestError)
	GetAllWorkout(token string, page int, limit int) (*WorkoutsResponse, *RateLimit, *RequestError)
	DeAuthorize(token string) (*RateLimit, *RequestError)
	GetUser(token string) (*User, *RateLimit, *RequestError)
}
