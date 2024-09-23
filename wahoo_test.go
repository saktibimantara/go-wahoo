package go_wahoo

import (
	"fmt"
	gohttp "github.com/saktibimantara/go-http"
	"net/url"
	"testing"
)

func TestWahoo_getRedirectParam(t *testing.T) {
	type fields struct {
		baseURL      string
		redirectURL  string
		clientID     string
		clientSecret string
		scopes       []OAuth2Scope
		goHTTP       *gohttp.GoHTTP
	}
	type args struct {
		uniqueCode string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test getRedirectParam",
			fields: fields{
				baseURL:      "https://api.wahooligan.com",
				redirectURL:  "https://abc.com",
				clientID:     "aaa",
				clientSecret: "bbb",
				scopes:       []OAuth2Scope{Email, UserRead, WorkoutsRead},
				goHTTP:       nil,
			},
			args: args{
				uniqueCode: "123",
			},
			want: "redirect_uri=" + url.PathEscape("https://abc.com?unique_code=123"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wahoo{
				baseURL:      tt.fields.baseURL,
				redirectURL:  tt.fields.redirectURL,
				clientID:     tt.fields.clientID,
				clientSecret: tt.fields.clientSecret,
				scopes:       tt.fields.scopes,
				goHTTP:       tt.fields.goHTTP,
			}
			if got := w.getRedirectParam(tt.args.uniqueCode); got != tt.want {
				t.Errorf("getRedirectParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWahoo_getScopeParam(t *testing.T) {
	type fields struct {
		baseURL      string
		redirectURL  string
		clientID     string
		clientSecret string
		scopes       []OAuth2Scope
		goHTTP       *gohttp.GoHTTP
	}
	type args struct {
		uniqueCode string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test getRedirectParam",
			fields: fields{
				baseURL:      "https://api.wahooligan.com",
				redirectURL:  "https://abc.com",
				clientID:     "aaa",
				clientSecret: "bbb",
				scopes:       []OAuth2Scope{Email, UserRead, WorkoutsRead},
				goHTTP:       nil,
			},
			args: args{
				uniqueCode: "123",
			},
			want: "scope=email+user_read+workouts_read",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wahoo{
				baseURL:      tt.fields.baseURL,
				redirectURL:  tt.fields.redirectURL,
				clientID:     tt.fields.clientID,
				clientSecret: tt.fields.clientSecret,
				scopes:       tt.fields.scopes,
				goHTTP:       tt.fields.goHTTP,
			}
			if got := w.getScopeParam(); got != tt.want {
				t.Errorf("getScopeParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWahoo(t *testing.T) {
	iface := NewWahoo("9F7GT9plmXV393jecMI5hpJg7953JLBPXTmSpQ5OFEI", "HQE-wMqpKws4b5lJy33EZzr8x_jguPs36UUefnqkmrU")
	iface.SetScopes(UserRead, WorkoutsRead, OfflineData, PowerZonesRead)
	iface.SetRedirectURI("https://champion-buck-pet.ngrok-free.app/api/v1/wahoo/exchange_token")

	//authenticateURL, err := iface.GetAuthenticateURL("123-onboarding")
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(*authenticateURL)

	//access, err := iface.GetAccessToken("UjONcvdRHeZ8XKODBYT2wFNr9s_3ST65-rTrMOlRBdw", "123-onboarding")
	//if err != nil {
	//	return
	//}
	//
	//fmt.Printf("Access: %v\n", *access)

	token, _, err := iface.RefreshToken("Qm74UB4_pbR8ZWHjl_ekjAjUEVgeNr910-j8oanc0mE", "123")
	if err != nil {
		return
	}

	fmt.Printf("Token: %v\n", *token)

	//{ESXjiRVX9CKBgEVwN917jJk7iMpVVNZBMh2xgrhJl94 Bearer 7199 jEue3OSJwF-obYKcVi-atmtCAnGqtir_4nH8FE5OrcA user_read workouts_read offline_data power_zones_read 1727064420}

}
