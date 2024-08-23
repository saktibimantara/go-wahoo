package go_wahoo

import (
	"fmt"
	"strings"

	gohttp "github.com/saktibimantara/go-http"
)

type OAuth2Scope string

const (
	Email         OAuth2Scope = "email"
	UserRead      OAuth2Scope = "user_read"
	UserWrite     OAuth2Scope = "user_write"
	WorkoutsRead  OAuth2Scope = "workouts_read"
	WorkoutsWrite OAuth2Scope = "workouts_write"
	OfflineData   OAuth2Scope = "offline_data"
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
	authenticateURL := fmt.Sprintf("%s/oauth/authorize?%s&%s&%s?unique_code=%s&response_type=code", w.baseURL, w.getClientParams(), w.getScopeParam(), w.getRedirectParam(), uniqueCode)

	return &authenticateURL, nil
}

func (w *Wahoo) GetAccessToken(code string) (*TokenResponse, *RequestError) {
	if err := w.validateAccessTokenRequest(code); err != nil {
		return nil, NewError(err, 400, "Invalid Request")
	}

	// buildAccessTokenURL
	accessTokenURL := fmt.Sprintf("%s/oauth/token?%s&%s&grant_type=authorization_code&code=%s", w.baseURL, w.getClientParams(), w.getRedirectParam(), code)

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

func (w *Wahoo) RefreshToken(refreshToken string) (*TokenResponse, *RequestError) {
	if err := w.validateRefreshTokenRequest(refreshToken); err != nil {
		return nil, NewError(err, 400, "Invalid Request")
	}

	// buildAccessTokenURL
	refreshTokenURL := fmt.Sprintf("%s/oauth/token?%s&%s&grant_type=refresh_token&refresh_token=%s", w.baseURL, w.getClientParams(), w.getRedirectParam(), refreshToken)

	// request to get access token
	resp, err := w.goHTTP.Post(refreshTokenURL, nil)
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

func (w *Wahoo) getClientParams() string {
	return fmt.Sprintf("client_id=%s&client_secret=%s", w.clientID, w.clientSecret)
}

func (w *Wahoo) getRedirectParam() string {
	return "redirect_uri=" + w.redirectURL
}

func (w *Wahoo) getScopeParam() string {
	scopes := "scopes="

	if len(w.scopes) == 0 {
		return ""
	}

	for i, scope := range w.scopes {
		if i == 0 {
			scopes += string(scope)
		} else {
			scopes += " " + string(scope)
		}
	}

	// replace space with %20 since URL.ParseQuery will parse space as +
	scopes = strings.ReplaceAll(scopes, " ", "%20")

	return scopes
}

func (w *Wahoo) GetAllWorkout(token string, page int, limit int) (*WorkoutsResponse, *RequestError) {
	workoutsURL := fmt.Sprintf("v1/workouts?page=%d&limit=%d", page, limit)

	w.SetBearerToken(token)

	resp, err := w.goHTTP.Get(workoutsURL)
	if err != nil {
		return nil, NewError(err, 500, "failed to get all workout")
	}

	if resp.Code != 200 {
		return nil, NewError(ErrGetAllWorkout, resp.Code, string(resp.Data))
	}

	return UnmarshalToWorkoutsResponse(resp.Data)
}
