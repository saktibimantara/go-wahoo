package go_wahoo

type IWahoo interface {
	GetAuthenticateURL(uniqueCode string) (*string, error)
	GetAccessToken(code string) (*TokenResponse, *RequestError)
	RefreshToken(refreshToken string) (*TokenResponse, *RequestError)
	GetAllWorkout(token string, page int, limit int) (*WorkoutsResponse, *RequestError)
}
