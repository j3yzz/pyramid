package database

import "gorm.io/gorm"

type Connection interface {
	GetConnection() (*gorm.DB, error)
}
