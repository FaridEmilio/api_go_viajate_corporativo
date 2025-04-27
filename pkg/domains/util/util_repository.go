package util

import (
	"github.com/faridEmilio/api_go_viajate_corporativo/internal/database"
)

type UtilRepository interface {
}

func NewUtilRepository(conn *database.MySQLClient) UtilRepository {
	return &utilRepository{
		SqlClient: conn,
	}
}

type utilRepository struct {
	SqlClient *database.MySQLClient
}
