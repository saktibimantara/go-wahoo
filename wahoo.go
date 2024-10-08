package go_wahoo

import (
	"fmt"
	"net/url"

	gohttp "github.com/saktibimantara/go-http"
)

type OAuth2Scope string

const (
	Email           OAuth2Scope = "email"
	UserRead        OAuth2Scope = "user_read"
	UserWrite       OAuth2Scope = "user_write"
	WorkoutsRead    OAuth2Scope = "workouts_read"
	WorkoutsWrite   OAuth2Scope = "workouts_write"
	OfflineData     OAuth2Scope = "offline_data"
	PowerZonesRead  OAuth2Scope = "power_zones_read"
	PowerZonesWrite OAuth2Scope = "power_zones_write"
	PlansRead       OAuth2Scope = "plans_read"
	PlansWrite      OAuth2Scope = "plans_write"
)

// Wahoo represents a struct with a baseURL field
type Wahoo struct {
	baseURL      string
	redirectURL  string
	clientID     string
	clientSecret string
	scopes       []OAuth2Scope
	goHTTP       *gohttp.GoHTTP
}

// NewWahoo creates a new Wahoo instance with a default baseURL
func NewWahoo(clientID, clientSecret string) *Wahoo {
	wh := Wahoo{
		baseURL:      "https://api.wahooligan.com",
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  "",
	}

	header := gohttp.Header{}
	header["Content-Type"] = "application/json"

	goHTTP := gohttp.New(&gohttp.Config{BaseURL: wh.baseURL, Header: header})
	wh.goHTTP = goHTTP

	return &wh
}

func (w *Wahoo) SetBearerToken(token string) *Wahoo {
	header := gohttp.NewBearerToken(token)
	header["Content-Type"] = "application/json"

	goHTTP := gohttp.New(&gohttp.Config{BaseURL: w.baseURL, Header: header})
	w.goHTTP = goHTTP

	return w
}

func (w *Wahoo) SetRedirectURI(uri string) *Wahoo {
	w.redirectURL = uri

	return w
}

func (w *Wahoo) SetScopes(scopes ...OAuth2Scope) *Wahoo {
	w.scopes = scopes

	return w
}

// GetAuthenticateURL UniqueCode is a string to represent who is the requester
func (w *Wahoo) GetAuthenticateURL(uniqueCode string) (*string, error) {
	if err := w.validateAuthenticate(); err != nil {
		return nil, err
	}

	// buildAuthenticateURL
	authenticateURL := fmt.Sprintf("%s/oauth/authorize?%s&%s&%s&response_type=code", w.baseURL, w.getClientParams(), w.getScopeParam(), w.getRedirectParam(uniqueCode))

	return &authenticateURL, nil
}

func (w *Wahoo) GetAccessToken(code, uniqueCode string) (*TokenResponse, *RequestError) {
	if err := w.validateAccessTokenRequest(code); err != nil {
		return nil, NewError(err, 400, "Invalid Request")
	}

	// buildAccessTokenURL
	accessTokenURL := fmt.Sprintf("%s/oauth/token?%s&%s&grant_type=authorization_code&code=%s&scopes=%s", w.baseURL, w.getAuthorizationClientParam(), w.getRedirectParam(uniqueCode), code, w.getScopeParam())

	// request to get access token
	resp, err := w.goHTTP.Post(accessTokenURL, nil)
	if err != nil {
		return nil, NewError(err, 500, "failed to get access token")
	}

	respMessage := string(resp.Data)

	if resp.Code != 200 {
		return nil, NewError(ErrFailedToGetAccessToken, resp.Code, respMessage)
	}

	if resp.Data == nil {
		return nil, NewError(ErrFailedToGetAccessToken, 500, "failed to get access token")
	}

	return UnmarshalToResponse(resp.Data)
}

func (w *Wahoo) RefreshToken(refreshToken, uniqueCode string) (*TokenResponse, *RateLimit, *RequestError) {
	if err := w.validateRefreshTokenRequest(refreshToken); err != nil {
		return nil, nil, NewError(err, 400, "Invalid Request")
	}

	// buildAccessTokenURL
	refreshTokenURL := fmt.Sprintf("%s/oauth/token?%s&grant_type=refresh_token&refresh_token=%s", w.baseURL, w.getAuthorizationClientParam(), refreshToken)

	// request to get access token
	resp, err := w.goHTTP.Post(refreshTokenURL, nil)
	if err != nil {
		return nil, nil, NewError(err, 500, "failed to get access token")
	}

	respMessage := string(resp.Data)

	rateLimit := NewRateLimit(resp.Header)

	if resp.Code != 200 {
		return nil, rateLimit, NewError(ErrFailedToGetAccessToken, resp.Code, respMessage)
	}

	if resp.Data == nil {
		return nil, rateLimit, NewError(ErrFailedToGetAccessToken, 500, "failed to get access token")
	}

	var token TokenResponse

	errUnmarshal := UnmarshalResponse(&token, resp.Data)
	return &token, rateLimit, errUnmarshal
}

func (w *Wahoo) GetUser(token string) (*User, *RateLimit, *RequestError) {
	userURL := "/v1/user"

	w.SetBearerToken(token)

	resp, err := w.goHTTP.Get(userURL)
	if err != nil {
		return nil, nil, NewError(err, 500, "failed to get user")
	}

	rateLimit := NewRateLimit(resp.Header)

	if resp.Code != 200 {
		return nil, rateLimit, NewError(ErrGetAllWorkout, resp.Code, string(resp.Data))
	}

	var user User
	errUnmarshal := UnmarshalResponse(&user, resp.Data)

	return &user, rateLimit, errUnmarshal
}

func (w *Wahoo) validateAccessTokenRequest(code string) error {
	if code == "" {
		return ErrInvalidCode
	}

	if w.clientID == "" {
		return ErrInvalidClientID
	}

	if w.clientSecret == "" {
		return ErrInvalidClientSecret
	}

	return nil
}

func (w *Wahoo) validateAuthenticate() error {
	if w.redirectURL == "" {
		return ErrInvalidRedirectURI
	}

	if len(w.scopes) == 0 {
		return ErrInvalidScopes
	}

	if w.clientID == "" {
		return ErrInvalidClientID
	}

	if w.clientSecret == "" {
		return ErrInvalidClientSecret
	}

	return nil
}

func (w *Wahoo) validateRefreshTokenRequest(refreshToken string) error {
	if refreshToken == "" {
		return ErrInvalidRefreshToken
	}

	if w.clientID == "" {
		return ErrInvalidClientID
	}

	if w.clientSecret == "" {
		return ErrInvalidClientSecret
	}

	return nil
}

func (w *Wahoo) getAuthorizationClientParam() string {
	return w.getClientParams() + "&" + w.getClientSecretParams()
}

func (w *Wahoo) getClientSecretParams() string {
	return "client_secret=" + w.clientSecret
}

func (w *Wahoo) getClientParams() string {
	return "client_id=" + w.clientID
}

func (w *Wahoo) getRedirectParam(uniqueCode string) string {
	urlStr := "redirect_uri=" + w.redirectURL + "?unique_code=" + uniqueCode

	// encode the url
	return url.PathEscape(urlStr)
}

func (w *Wahoo) getScopeParam() string {
	scopes := "scope="

	if len(w.scopes) == 0 {
		return ""
	}

	for i, scope := range w.scopes {
		if i == 0 {
			scopes += string(scope)
		} else {
			scopes += "+" + string(scope)
		}
	}

	return scopes
}

func (w *Wahoo) GetAllWorkout(token string, page int, limit int) (*WorkoutsResponse, *RateLimit, *RequestError) {
	workoutsURL := fmt.Sprintf("/v1/workouts?page=%d&limit=%d", page, limit)

	w.SetBearerToken(token)

	resp, err := w.goHTTP.Get(workoutsURL)
	if err != nil {
		return nil, nil, NewError(err, 500, "failed to get all workout")
	}

	rateLimit := NewRateLimit(resp.Header)

	if resp.Code != 200 {
		return nil, rateLimit, NewError(ErrGetAllWorkout, resp.Code, string(resp.Data))
	}

	var workouts WorkoutsResponse
	errUnmarshal := UnmarshalResponse(&workouts, resp.Data)

	return &workouts, rateLimit, errUnmarshal
}

func (w *Wahoo) DeAuthorize(token string) (*RateLimit, *RequestError) {
	deAuthorizeURL := w.baseURL + "/v1/permissions"

	w.SetBearerToken(token)

	resp, err := w.goHTTP.Delete(deAuthorizeURL)
	if err != nil {
		return nil, NewError(err, 500, "failed to deAuthorize")
	}

	rateLimit := NewRateLimit(resp.Header)

	if resp.Code != 200 {
		return rateLimit, NewError(ErrDeAuthorize, resp.Code, string(resp.Data))
	}

	return rateLimit, nil
}
