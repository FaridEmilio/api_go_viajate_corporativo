package comunidaddtos

type RequestTipoComunidad struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Activo bool   `json:"activo"`
	Size   int64  `json:"size"`
	Number int64  `json:"number"`
}
