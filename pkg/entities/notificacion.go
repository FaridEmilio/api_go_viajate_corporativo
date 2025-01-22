package entities

import (
	"errors"

	"gorm.io/gorm"
)

type Notificacione struct {
	gorm.Model
	UserId      uint64               `json:"user_id"`
	Tipo        EnumTipoNotificacion `json:"tipo"`
	Descripcion string               `json:"descripcion"`
}

type EnumTipoNotificacion string

const (
	NotificacionTransferencia   EnumTipoNotificacion = "Transferencia"
	NotificacionCierreLote      EnumTipoNotificacion = "CierreLote"
	NotificacionPagoExpirado    EnumTipoNotificacion = "PagoExpirado"
	NotificacionConfiguraciones EnumTipoNotificacion = "Configuraciones"
	NotificacionSolicitudCuenta EnumTipoNotificacion = "SolicitudCuenta"
	NotivicacionEnvioEmail      EnumTipoNotificacion = "EnvioEmail"
)

func (e EnumTipoNotificacion) IsValid() error {
	switch e {
	case NotificacionTransferencia, NotificacionCierreLote:
		return nil
	}
	return errors.New("tipo EnumTipoNotificacion con formato inválido")
}
