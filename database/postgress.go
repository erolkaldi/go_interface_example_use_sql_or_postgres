package database

import (
	"cmd/main.go/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	ConnectionString string
	Dialector        gorm.Dialector
}

func (pgs Postgress) Connect(config *models.Config) (*gorm.DB, error) {
	pgs.ConnectionString = fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=%s", config.Postgre.Host, config.Postgre.Port, config.Postgre.Password, config.Postgre.DbName, config.Postgre.SSLMode)
	pgs.Dialector = postgres.Open(pgs.ConnectionString)
	return gorm.Open(pgs.Dialector, &gorm.Config{})
}
