package authdtos

type ResponseLogin struct {
	User         ResponseUsuario `json:"usuario"`
	Token        string          `json:"token"`
	RefreshToken string          `json:"refresh_token"`
}
