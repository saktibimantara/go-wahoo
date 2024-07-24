package go_wahoo

import (
	"reflect"
	"testing"
)

func TestUnmarshalToResponse(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    *TokenResponse
		wantErr bool
	}{
		{
			name: "Test UnmarshalToResponse",
			args: args{
				data: `{
  "access_token": "9IGrKxQKfhwld32SFv9nCRT3jptoAmshINrFEpQZ7Kw",
  "token_type": "Bearer",
  "expires_in": 7199,
  "refresh_token": "yOXxKK2p90C1H5P0EKuBciv3vNesptYMfGzUwTR5MMg",
  "scope": "user_read",
  "created_at": 1721808795
}`,
			},
			want: &TokenResponse{
				AccessToken:  "9IGrKxQKfhwld32SFv9nCRT3jptoAmshINrFEpQZ7Kw",
				TokenType:    "Bearer",
				ExpiresIn:    7199,
				RefreshToken: "yOXxKK2p90C1H5P0EKuBciv3vNesptYMfGzUwTR5MMg",
				Scope:        "user_read",
				CreatedAt:    1721808795,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []byte(tt.args.data)

			got, err := UnmarshalToResponse(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalToResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalToResponse() got = %v, want %v", got, tt.want)
			}

			if got.GetAccessToken() != tt.want.GetAccessToken() {
				t.Errorf("UnmarshalToResponse() got = %v, want %v", got.GetAccessToken(), tt.want.GetAccessToken())
			}

			if got.GetTokenType() != tt.want.GetTokenType() {
				t.Errorf("UnmarshalToResponse() got = %v, want %v", got.GetTokenType(), tt.want.GetTokenType())
			}

			if got.GetExpiresAt() != tt.want.GetExpiresAt() {
				t.Errorf("UnmarshalToResponse() got = %v, want %v", got.GetExpiresAt(), tt.want.GetExpiresAt())
			}

			if got.GetRefreshToken() != tt.want.GetRefreshToken() {
				t.Errorf("UnmarshalToResponse() got = %v, want %v", got.GetRefreshToken(), tt.want.GetRefreshToken())
			}

			if got.GetScope() != tt.want.GetScope() {
				t.Errorf("UnmarshalToResponse() got = %v, want %v", got.GetScope(), tt.want.GetScope())
			}
		})
	}
}
