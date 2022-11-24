package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Mysql struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
}

func (m Mysql) GetConnection() (*gorm.DB, error) {
	host, ok := os.LookupEnv("MYSQL_HOST")
	logEnvNotFound(ok, "MYSQL_HOST")
	m.Host = host

	username, ok := os.LookupEnv("MYSQL_USERNAME")
	logEnvNotFound(ok, "MYSQL_USERNAME")
	m.Username = username

	password, ok := os.LookupEnv("MYSQL_PASSWORD")
	logEnvNotFound(ok, "MYSQL_PASSWORD")
	m.Password = password

	database, ok := os.LookupEnv("MYSQL_DATABASE")
	logEnvNotFound(ok, "MYSQL_DATABASE")
	m.Database = database

	port, ok := os.LookupEnv("MYSQL_PORT")
	logEnvNotFound(ok, "MYSQL_PORT")
	m.Port = port

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mysql Connected.")
	return connection, nil
}

func logEnvNotFound(ok bool, key string) {
	if !ok {
		log.Fatal(key + " Not found in .env")
	}
}
