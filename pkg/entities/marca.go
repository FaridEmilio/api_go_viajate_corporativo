package entities

// Marca representa una marca de vehículo (Toyota, Ford, etc.)
type Marca struct {
	ID    uint
	Marca string `json:"marca"`
}
