package main

import (
	"github.com/j3yzz/pyramid/infrastructure/database"
	"github.com/j3yzz/pyramid/infrastructure/database/driver"
	"github.com/j3yzz/pyramid/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	envErr := godotenv.Load("./.env")
	if envErr != nil {
		log.Fatal("Failed to loading .env file")
	}

	migration := database.Migrations{DB: driver.Mysql{}}
	migration.MakeMigrations()

	router.InitRoutes()

	server := &http.Server{
		Handler:      router.GetRouter(),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
