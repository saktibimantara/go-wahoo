package go_wahoo

import (
	"fmt"
	"testing"
)

func TestWahoo_GetAuthenticateURL(t *testing.T) {

	iface := NewWahoo("aaa", "bbb").SetScopes(Email, UserRead, WorkoutsRead).SetRedirectURI("ccc.com")

	tests := []struct {
		name    string
		wahoo   IWahoo
		want    *string
		wantURL string
	}{
		{
			name:    "Test Case 1",
			want:    nil,
			wahoo:   iface,
			wantURL: "https://api.wahooligan.com/oauth/authorize?client_id=aaa&client_secret=bbb&scopes=email%20user_read%20workouts_read&redirect_uri=ccc.com&response_type=code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.wahoo.GetAuthenticateURL()
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

	c := NewWahoo("aaa", "bbb").SetRedirectURI("ccc.com").SetScopes(Email, UserRead, WorkoutsRead)

	token, err := c.GetAccessToken("V9eugS4YlZhShFyRZSSGjtRsGSyPzcGGZHRB2Ip4HIA")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if token == nil {
		t.Error("Token is nil")
		t.Fail()
		return
	}

	fmt.Println("AccessToken", token.GetAccessToken())
	fmt.Println("RefreshToken", token.GetRefreshToken())
	fmt.Println("ExpiresIn", token.GetExpiresAt())
	fmt.Println("Scope", token.GetScope())
	fmt.Println("CreatedAt", token.GetCreatedAt())
}
