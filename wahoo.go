package go_wahoo

import (
	"errors"
	"fmt"
	go_http "github.com/saktibimantara/go-http"
	"strings"
)

var (
	AuthorizeURL = "%s/oauth/authorize"
	TokenURL     = "%s/oauth/token"
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
	clientId     string
	clientSecret string
	scopes       []OAuth2Scope
	goHttp       *go_http.GoHTTP
}

// NewWahoo creates a new Wahoo instance with a default baseURL
func NewWahoo(clientId, clientSecret string) *Wahoo {

	wh := Wahoo{
		baseURL:      "https://api.wahooligan.com",
		clientId:     clientId,
		clientSecret: clientSecret,
	}

	goHttp := go_http.New(&go_http.Config{BaseURL: wh.baseURL})
	wh.goHttp = goHttp

	return &wh
}

func (w *Wahoo) SetRedirectURI(uri string) *Wahoo {
	w.redirectURL = uri

	return w
}

func (w *Wahoo) SetScopes(scopes ...OAuth2Scope) *Wahoo {
	w.scopes = scopes

	return w
}

type IWahoo interface {
	GetAuthenticateURL() (*string, error)
}

func (w *Wahoo) GetAuthenticateURL() (*string, error) {
	if err := w.validateAuthenticate(); err != nil {
		return nil, err
	}

	// buildAuthenticateURL
	authenticateURL := fmt.Sprintf("%s/oauth/authorize?%s&%s&%s&response_type=code", w.baseURL, w.getClientParams(), w.getScopeParam(), w.getRedirectParam())

	return &authenticateURL, nil
}

func (w *Wahoo) validateAuthenticate() error {
	if w.redirectURL == "" {
		return errors.New("redirect url is required")
	}

	if len(w.scopes) == 0 {
		return errors.New("scopes is required")
	}

	if w.clientId == "" {
		return errors.New("client id is required")
	}

	if w.clientSecret == "" {
		return errors.New("client secret is required")
	}

	return nil
}

func (w *Wahoo) getClientParams() string {
	return fmt.Sprintf("client_id=%s&client_secret=%s", w.clientId, w.clientSecret)
}

func (w *Wahoo) getRedirectParam() string {
	return fmt.Sprintf("redirect_uri=%s", w.redirectURL)
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
			scopes += fmt.Sprintf(" %s", string(scope))
		}
	}

	// replace space with %20 since URL.ParseQuery will parse space as +
	scopes = strings.ReplaceAll(scopes, " ", "%20")

	return scopes
}
