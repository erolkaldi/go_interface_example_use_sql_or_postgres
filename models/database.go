package models

import (
	"gorm.io/gorm"
)

type BaseDb interface {
	Connect(config *Config) (*gorm.DB, error)
}
