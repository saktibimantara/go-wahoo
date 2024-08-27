package tests

import (
	"fmt"
	"github.com/saktibimantara/go-wahoo"
	"github.com/saktibimantara/go-wahoo/mocks"
	"testing"
)

func TestWahoo_GetAuthenticateURL(t *testing.T) {

	iface := go_wahoo.NewWahoo("aaa", "bbb").SetScopes(go_wahoo.Email, go_wahoo.UserRead, go_wahoo.WorkoutsRead).SetRedirectURI("ccc.com")

	tests := []struct {
		name    string
		wahoo   go_wahoo.IWahoo
		want    *string
		wantURL string
	}{
		{
			name:    "Test Case 1",
			want:    nil,
			wahoo:   iface,
			wantURL: "https://api.wahooligan.com/oauth/authorize?client_id=aaa&scope=email+user_read+workouts_read&redirect_uri=ccc.com%3Funique_code=123&response_type=code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.wahoo.GetAuthenticateURL("123")
			if err != nil {
				t.Errorf("Wahoo.GetAuthenticateURL() error = %v, wantErr %v", err, tt.want)
				return
			}

			if *got != tt.wantURL {
				t.Errorf("Wahoo.GetAuthenticateURL() = %v, want %v", *got, tt.wantURL)
			}
		})
	}
}

func TestWahoo_GetAccessToken(t *testing.T) {

	w := mocks.IWahoo{}

	w.On("GetAccessToken", "abcdfsd", "123").Return(&go_wahoo.TokenResponse{
		AccessToken:  "9IGrKxQKfhwld32SFv9nCRT3jptoAmshINrFEpQZ7Kw",
		TokenType:    "Bearer",
		ExpiresIn:    7199,
		RefreshToken: "yOXxKK2p90C1H5P0EKuBciv3vNesptYMfGzUwTR5MMg",
		Scope:        "user_read",
		CreatedAt:    1721808795,
	}, nil)

	w.On("GetAccessToken", "badCode").Return(nil, &go_wahoo.RequestError{
		Err:   go_wahoo.ErrFailedToGetAccessToken,
		Code:  400,
		Debug: "Failed to get access token",
	})

	token, err := w.GetAccessToken("abcdfsd", "123")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if token == nil {
		t.Error("Token is nil")
		t.Fail()
		return
	}

	if token.GetAccessToken() != "9IGrKxQKfhwld32SFv9nCRT3jptoAmshINrFEpQZ7Kw" {
		t.Error("Access token is not valid")
		t.Fail()
		return
	}
}

func TestWahoo_GetRefreshToken(t *testing.T) {

	w := mocks.IWahoo{}

	w.On("RefreshToken", "refresh_001", "1234").Return(&go_wahoo.TokenResponse{
		AccessToken:  "9IGrKxQKfhwld32SFv9nCRT3jptoAmshINrFEpQZ7Kw",
		TokenType:    "Bearer",
		ExpiresIn:    7199,
		RefreshToken: "yOXxKK2p90C1H5P0EKuBciv3vNesptYMfGzUwTR5MMg",
		Scope:        "user_read",
		CreatedAt:    1721808795,
	}, nil)

	w.On("RefreshToken", "badRefreshToken").Return(nil, &go_wahoo.RequestError{
		Err: go_wahoo.ErrFailedToGetAccessToken,
	})

	refreshToken, err := w.RefreshToken("refresh_001", "1234")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if refreshToken == nil {
		t.Error("Token is nil")
		t.Fail()
		return
	}

	if refreshToken.GetAccessToken() != "9IGrKxQKfhwld32SFv9nCRT3jptoAmshINrFEpQZ7Kw" {
		t.Error("Access token is not valid")
		t.Fail()
		return
	}

	if refreshToken.GetRefreshToken() !=
		"yOXxKK2p90C1H5P0EKuBciv3vNesptYMfGzUwTR5MMg" {
		t.Error("Refresh token is not valid")
		t.Fail()
		return
	}

	if refreshToken.GetScope() != "user_read" {
		t.Error("Scope is not valid")
		t.Fail()
		return
	}

}
