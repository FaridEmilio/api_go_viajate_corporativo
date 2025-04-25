package authdtos

type VerifyEmailRequest struct {
	Token string `json:"token"`
}

type ResendVerifyEmailRequest struct {
	Email string `json:"email"`
}
