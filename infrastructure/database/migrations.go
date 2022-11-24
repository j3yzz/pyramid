package database

import (
	"github.com/j3yzz/pyramid/app/users/migrations"
	"log"
)

type Migrations struct {
	DB Connection
}

func (m Migrations) MakeMigrations() {
	db, err := m.DB.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("migrating")

	migrations.CreateUsersTable(db)
}
