package entities

type Estado struct {
	ID     uint
	Nombre string `json:"nombre"`
}
type EnumEstado string

const (
	Disponible EnumEstado = "Disponible"
	Expirado   EnumEstado = "Expirado"
	Cancelado  EnumEstado = "Cancelado"
	Completo   EnumEstado = "Completo"
	Pendiente  EnumEstado = "Pendiente"
	Vencida    EnumEstado = "Vencida"
	Aceptada   EnumEstado = "Aceptada"
	Rechazada  EnumEstado = "Rechazada"
)
