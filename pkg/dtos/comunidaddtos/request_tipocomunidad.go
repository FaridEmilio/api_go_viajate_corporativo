package comunidaddtos

type RequestTipoComunidad struct {
	Id     int    `json:"id"`
	Tipo   string `json:"Tipo"`
	Activo bool   `json:"activo"`
	Size   int64  `json:"size"`
	Number int64  `json:"number"`
}
