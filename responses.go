package go_wahoo

import (
	"encoding/json"
	"time"
)

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"createdAt"`
}

type WorkoutsResponse struct {
	Workouts []Workout `json:"workouts"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PerPage  int       `json:"per_page"`
	Order    string    `json:"order"`
	Sort     string    `json:"sort"`
}

type Workout struct {
	Id             int            `json:"id"`
	Starts         time.Time      `json:"starts"`
	Minutes        int            `json:"minutes"`
	Name           string         `json:"name"`
	PlanId         *int           `json:"plan_id"`
	WorkoutToken   string         `json:"workout_token"`
	WorkoutTypeId  int            `json:"workout_type_id"`
	WorkoutSummary WorkoutSummary `json:"workout_summary"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type WorkoutSummary struct {
	Id                  int         `json:"id"`
	AscentAccum         string      `json:"ascent_accum"`
	CadenceAvg          string      `json:"cadence_avg"`
	CaloriesAccum       string      `json:"calories_accum"`
	DistanceAccum       string      `json:"distance_accum"`
	DurationActiveAccum string      `json:"duration_active_accum"`
	DurationPausedAccum string      `json:"duration_paused_accum"`
	DurationTotalAccum  string      `json:"duration_total_accum"`
	HeartRateAvg        string      `json:"heart_rate_avg"`
	PowerBikeNpLast     string      `json:"power_bike_np_last"`
	PowerBikeTssLast    string      `json:"power_bike_tss_last"`
	PowerAvg            string      `json:"power_avg"`
	SpeedAvg            string      `json:"speed_avg"`
	WorkAccum           string      `json:"work_accum"`
	File                WorkoutFile `json:"file"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
}

type WorkoutFile struct {
	Url string `json:"url"`
}

type ITokenResponse interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() time.Time
	GetRefreshToken() string
	GetScope() string
	GetCreatedAt() time.Time
}

func UnmarshalToWorkoutsResponse(data []byte) (*WorkoutsResponse, *RequestError) {
	var resp WorkoutsResponse
	err := json.Unmarshal(data, &resp)

	if err != nil {
		return nil, NewError(err, 500, "failed to unmarshal response")
	}

	return &resp, nil
}

func UnmarshalToResponse(data []byte) (*TokenResponse, *RequestError) {
	var resp TokenResponse
	err := json.Unmarshal(data, &resp)

	if err != nil {
		return nil, NewError(err, 500, "failed to unmarshal response")
	}

	return &resp, nil
}

func (t TokenResponse) GetAccessToken() string {
	return t.AccessToken
}

func (t TokenResponse) GetTokenType() string {
	return t.TokenType
}

func (t TokenResponse) GetExpiresAt() time.Time {
	if t.ExpiresIn == 0 {
		return time.Time{}
	}

	return time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)
}

func (t TokenResponse) GetRefreshToken() string {
	return t.RefreshToken
}

func (t TokenResponse) GetScope() string {
	return t.Scope
}

func (t TokenResponse) GetCreatedAt() time.Time {
	return time.Unix(t.CreatedAt, 0)
}
