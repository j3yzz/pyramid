package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router

func GetRouter() *mux.Router {
	return router
}

func Init() {
	router = mux.NewRouter()
}

func InitRoutes() {
	Init()
	router.Headers("Content-Type", "application/json")
	r := router.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]string{
			"message": "404 Not found",
		})
	}).Methods("GET")
	
}
