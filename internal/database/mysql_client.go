package database

import (
	"fmt"
	"time"

	"github.com/faridEmilio/api_go_gym_manager/internal/config"
	"github.com/faridEmilio/api_go_gym_manager/internal/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLClient contiene la instancia de base de datos
type MySQLClient struct {
	*gorm.DB
	TX *gorm.DB
}

// NewMySQLClient cliente de la base de datos en MySql
func NewMySQLClient() *MySQLClient {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.USER, config.PASSW, config.HOST, config.PORT, config.DB)
	// logs.Info(dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
			return time.Now().In(loc)
		},
	})

	database, _ := gormDB.DB()
	database.SetMaxIdleConns(20)
	database.SetMaxOpenConns(200)

	if err != nil {
		logs.Error("no se puede conectar la base de datos " + err.Error())
		panic(err)
	}

	return &MySQLClient{gormDB, nil}
}
