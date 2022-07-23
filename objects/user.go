package objects

type UserData struct {
	Tokens
	User User `json:"user"`
}

type AccessToken struct {
	Token          string `json:"token"`
	ExpirationDate string `json:"expiration_date"`
}

type User struct {
	ID       int    `json:"id"`
	Language string `json:"language"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Username string `json:"username"`
}
