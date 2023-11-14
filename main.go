package main

import (
	"virtualmachine/middlewares"
	"virtualmachine/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"virtualmachine/controllers/authcontroller"
	"virtualmachine/controllers/employeecontroller"
)

func main() {
	models.ConnectDB()
	r := mux.NewRouter()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/employees", employeecontroller.Index).Methods("GET")
	api.HandleFunc("/employee/{id}", employeecontroller.Show).Methods("GET")
	api.HandleFunc("/employee", employeecontroller.Create).Methods("POST")
	api.HandleFunc("/employee/{id}", employeecontroller.Update).Methods("PUT")
	api.HandleFunc("/employee", employeecontroller.Delete).Methods("DELETE")

	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
