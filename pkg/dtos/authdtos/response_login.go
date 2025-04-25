package authdtos

type ResponseLogin struct {
	User         ResponseUser `json:"user"`
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
}
