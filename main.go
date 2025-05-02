package main

import (
	"f1fastapi/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//маршруты
	r.HandleFunc("/drivers", handlers.GetAllDrivers).Methods("GET")
	r.HandleFunc("/drivers/{id}", handlers.GetDriverByID).Methods("GET")
	r.HandleFunc("/drivers", handlers.CreateDriver).Methods("POST")
	r.HandleFunc("/drivers/{id}", handlers.DeleteDriver).Methods("DELETE")
	//запуск сервера
	log.Println("Server running on http://localhost:6060")
	log.Fatal(http.ListenAndServe(":6060", r))
}
