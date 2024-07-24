package go_wahoo

import (
	"testing"
)

func TestWahoo_GetAuthenticateURL(t *testing.T) {

	iface := NewWahoo("wwwwwww", "aaa").SetScopes(Email, UserRead, WorkoutsRead).SetRedirectURI("http://localhost:8080")

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
			wantURL: "https://api.wahooligan.com?client_id=wwwwwww&client_secret=aaa&scopes=email%20user_read%20workouts_read&redirect_uri=http://localhost:8080",
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
