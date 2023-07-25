package database

import (
	"cmd/main.go/models"
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type SqlServer struct {
	ConnectionString string
	Dialector        gorm.Dialector
}

func (sql SqlServer) Connect(config *models.Config) (*gorm.DB, error) {
	sql.ConnectionString = fmt.Sprintf("server=" + config.SqlServer.Server + ";user id=" + config.SqlServer.User + ";password=" + config.SqlServer.Password + ";database=" + config.SqlServer.DbName + ";")
	sql.Dialector = sqlserver.Open(sql.ConnectionString)
	return gorm.Open(sql.Dialector, &gorm.Config{})
}
