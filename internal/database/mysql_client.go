package database

import (
	"fmt"
	"os"
	"time"

	"github.com/faridEmilio/api_go_viajate_corporativo/internal/logs"
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
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, pass, host, port, dbname)
	//logs.Info(dsn)
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
