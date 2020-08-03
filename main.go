package main

import (
	"authorizater/router"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	log.Printf("Server started")
	r := router.Router()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
