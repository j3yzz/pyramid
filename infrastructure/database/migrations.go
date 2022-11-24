package database

import "log"

type Migrations struct {
	DB Connection
}

func (m Migrations) MakeMigrations() {
	_, err := m.DB.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("migrating")

}
