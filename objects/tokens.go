package objects

type Tokens struct {
	AccessToken  AccessToken `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
