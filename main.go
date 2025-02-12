package main

import (
	"log"
	"net/http"

	handlers "github.com/Guaz-Delivery/guaz_backend/handlers"
	mux "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/signup_courier/", handlers.HandleCourierSignup)
	r.HandleFunc("/login_courier/", handlers.HandleCourierLogin)
	r.HandleFunc("/login_admin/", handlers.HandleAdminLogin)
	http.ListenAndServe(":9999", r)

}
