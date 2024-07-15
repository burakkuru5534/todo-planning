package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Veritabanı bağlantısını başlat
	database.InitDB()

	// Router oluştur
	r := mux.NewRouter()

	// Endpoint'leri tanımla
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/fetch-tasks", handlers.FetchTasks).Methods("POST")

	// HTTP sunucusunu başlat
	log.Fatal(http.ListenAndServe(":8080", r))
}
