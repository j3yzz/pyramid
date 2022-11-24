package migrations

import (
	"github.com/j3yzz/pyramid/app/users/models"
	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) {
	db.Migrator().AutoMigrate(&models.User{})
}
