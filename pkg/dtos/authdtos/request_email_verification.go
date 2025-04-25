package authdtos

type RequestUserEmailVerification struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
