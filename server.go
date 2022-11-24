package main

import (
	"github.com/j3yzz/pyramid/router"
	"log"
	"net/http"
	"time"
)

func main() {
	router.InitRoutes()

	server := &http.Server{
		Handler:      router.GetRouter(),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
