package entities

type Frecuencia struct {
	ID          uint
	Tipo        string
	Descripcion string
}

func (Frecuencia) TableName() string {
	return "frecuencias"
}
